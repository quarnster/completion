/*
 *
 * This file is in the public domain.
 *
 */
using System;
using System.IO;
using System.IO.Compression;

namespace PeterO
{
	/// <summary>
	/// A simple class for encoding PNG image files.
	/// </summary>
	public sealed class Png8BitIndexed
	{
		private byte[] Adler32(byte[] stream, int offset,
		                       int length){
			var adler = 1;
			var len = length;
			var NMAX = 3854;
			var BASE = 65521;
			var s1 = adler & 0xffff;
			var s2 = ((adler & 0xffff0000) >> 16)&0xFFFF;
			var k=0;
			var bpos = offset;
			while (len > 0) {
				k = len < NMAX ? len : NMAX;
				len -= k;
				while (k > 0) {
					s1 = unchecked((int)s1 + stream[bpos]);
					s2 = unchecked((int)s2 + s1);
					bpos += 1;
					k -= 1;
				}
				s1 = s1 % BASE;
				s2 = s2 % BASE;
			}
			return new byte[]{(byte)(s2>>8),
				(byte)(s2&255),
				(byte)(s1>>8),
				(byte)(s1&255)
			};
		}

		byte[] subdata1;
		uint[] crcTable;

		private uint Crc32(byte[] stream, int offset, int length, uint crc){
			uint c;
			if(crcTable==null){
				crcTable=new uint[256];
				for(uint n=0;n<=255;n++){
					c = n;
					for(var k=0;k<=7;k++){
						if((c & 1) == 1)
							c = 0xEDB88320^((c>>1)&0x7FFFFFFF);
						else
							c = ((c>>1)&0x7FFFFFFF);
					}
					crcTable[n] = c;
				}
			}
			c = crc^0xffffffff;
			var endOffset=offset+length;
			for(var i=offset;i<endOffset;i++){
				c = crcTable[(c^stream[i]) & 255]^((c>>8)&0xFFFFFF);
			}
			return c^0xffffffff;
		}

		/// Sets the PNG filter for the current row.
		public void SetFilter(int y, byte filter){
			this.data[y*this.realRowSize]=filter;
		}

		/// Gets the PNG filter for the current row.
		public byte GetFilter(int y){
			return this.data[y*this.realRowSize];
		}

		///
		///  Sets the pixel located at the X and Y coordinates of the
		///  image.  The pixel byte array is an index to the color table,
		/// 		from 1 though 15.
		///  Because the row may use a PNG filter, those components may
		///  not actually represent the index of each pixel returned.
		///
		public void SetPixel(int x,int y,int pixel){
			if(x<0||x>=width)throw new ArgumentOutOfRangeException("x");
			if(y<0||y>=height)throw new ArgumentOutOfRangeException("y");
			int offset=(y*this.realRowSize)+x+1;
			this.data[offset]=(byte)(pixel&0xFF);
		}
		///
		///  Gets the pixel located at the X and Y coordinates of the
		///  image.
		///
		public int GetPixel(int x,int y){
			if(x<0||x>=width)throw new ArgumentOutOfRangeException("x");
			if(y<0||y>=height)throw new ArgumentOutOfRangeException("y");
			int offset=(y*this.realRowSize)+x+1;
			return this.data[offset];
		}

		public byte[] GetColor(int index){
			if(index<0||index>=256)throw new ArgumentOutOfRangeException("index");
			return new byte[]{
				colors[index*3],
				colors[index*3+1],
				colors[index*3+2],255
			};
		}

		public void SetColor(int index, byte[] color){
			if(index<0||index>=256)throw new ArgumentOutOfRangeException("index");
			if(color==null)throw new ArgumentNullException("color");
			if(color.Length<3)throw new ArgumentException("'color' must have length 3 or more.");
			colors[index*3]=color[0];
			colors[index*3+1]=color[1];
			colors[index*3+2]=color[2];
		}

		private byte[] GetBE(uint crc){
			return new byte[]{
				(byte)((crc>>24)&255),
				(byte)((crc>>16)&255),
				(byte)((crc>>8)&255),
				(byte)((crc>>0)&255)
			};
		}

		/// Saves the image to a file.
		public void Save(string filename){
			using(FileStream fs=new FileStream(filename,FileMode.Create)){
				fs.Write(this.subdata1,0,this.subdata1.Length);
				uint crc32=Crc32(subdata1,12,17,0);
				byte[] deflated=null;
				fs.Write(GetBE(crc32),0,4);
				// Write the color data
				fs.Write(GetBE((uint)colors.Length),0,4);
				fs.Write(new byte[]{0x50,0x4c,0x54,0x45},0,4);
				crc32=Crc32(new byte[]{0x50,0x4c,0x54,0x45},0,4,0);
				crc32=Crc32(colors,0,colors.Length,crc32);
				fs.Write(colors,0,colors.Length);
				fs.Write(GetBE(crc32),0,4);
				// Write the transparent color
				if(transparent>=0 && transparent<256){
					fs.Write(GetBE(1),0,4);
					fs.Write(new byte[]{0x74,0x52,0x4e,0x53},0,4);
					crc32=Crc32(new byte[]{0x74,0x52,0x4e,0x53},0,4,0);
					crc32=Crc32(new byte[]{(byte)transparent},0,1,crc32);
					fs.Write(new byte[]{(byte)transparent},0,1);
					fs.Write(GetBE(crc32),0,4);
				}
				// Write the image data
				using(MemoryStream ms=new MemoryStream()){
					// PNG compression uses a ZLIB stream not a DEFLATE stream
					ms.WriteByte(0x78);
					ms.WriteByte(0x9c);
					using(DeflateStream ds=new DeflateStream(ms,
					                                         CompressionMode.Compress,true)){
						ds.Write(this.data,0,this.data.Length);
					}
					ms.Write(Adler32(this.data,0,this.data.Length),0,4);
					deflated=ms.ToArray();
				}
				byte[] defLength=new byte[]{
					(byte)((deflated.Length>>24)&255),
					(byte)((deflated.Length>>16)&255),
					(byte)((deflated.Length>>8)&255),
					(byte)((deflated.Length>>0)&255)
				};
				fs.Write(defLength,0,defLength.Length);
				fs.Write(new byte[]{
				         	0x49,0x44,0x41,0x54
				         },0,4);
				fs.Write(deflated,0,deflated.Length);
				uint crc=Crc32(deflated,0,deflated.Length,this.idatCrc);
				byte[] subdcrc=GetBE(crc);
				fs.Write(subdcrc,0,subdcrc.Length);
				fs.Write(subdata2,0,subdata2.Length);
			}
		}

		int width,height,realRowSize,blockSize,rowSize;


		/// Gets the height of the image.
		public int Height {
			get { return height; }
		}

		/// Gets the width of the image.
		public int Width {
			get { return width; }
		}
		byte[] imageData,data,subdata2,colors;
		uint idatCrc;
		int transparent;

		public int Transparent {
			get { return transparent; }
			set {
				if(value<0||value>=16)
					throw new ArgumentOutOfRangeException("value");
				transparent = value;
			}
		}
		/// Creates a new PNG image with the given
		/// width and height.
		public Png8BitIndexed(int width, int height)
		{
			if(width>65535 || width<=0)
				throw new ArgumentOutOfRangeException("width");
			if(height>65535 || height<=0)
				throw new ArgumentOutOfRangeException("height");
			colors=new byte[256*3];
			transparent=-1;
			subdata1=new byte[]{
				0x89,0x50,0x4e,0x47,0x0d,0x0a,0x1a,0x0a,0,0,0,0xd,
				0x49,0x48,0x44,0x52,
				0,0,(byte)(width>>8),(byte)(width&255),
				0,0,(byte)(height>>8),(byte)(height&255),
				8,3,0,0,0
			};
			this.width=width;
			this.height=height;
			this.realRowSize=this.width+1;
			this.blockSize=this.realRowSize;
			this.rowSize=this.realRowSize;
			this.imageData=new byte[this.rowSize*this.height];
			this.data=this.imageData;
			idatCrc=Crc32(
				new byte[]{
					0x49,0x44,0x41,0x54
				},0,4,0);
			subdata2=new byte[]{
				0,0,0,0,0x49,0x45,0x4e,0x44,0xae,0x42,0x60,0x82
			};
		}
	}
}
