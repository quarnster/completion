// Written by Peter O. in 2011.
// This file is in the public domain.

// Define to enable storing QR Codes in lightweight Png4BitIndexed objects
// (requires the Png4BitIndexed class at
//  http://articles.upokecenter.com/2011/05/11/png-4-bit-and-8-bit-indexed-encoder-classes-in-c/
//#define QRCODEPNG
// Define to enable storing QR Codes in System.Drawing.Bitmap objects
// (requires a reference to System.Drawing)
#define QRCODESYSTEMDRAWING

using System;
using System.Collections.Generic;

namespace PeterO
{
	/// <summary>
	/// A public domain static class for generating QR Code images.
	/// QR Code is becoming increasingly popular for presenting information
	/// in a printable, yet machine-readable form.
	/// QR Code supports various sizes and error correction levels, and this
	/// class generates the smallest size QR Code that can hold the specified
	/// amount of data, with error correction level M (medium, can correct up
	/// to 15% of data errors).
	///  Usage example (requires a reference to System.Drawing and defining
	///  QRCODESYSTEMDRAWING):
	/// <code>
	/// // Generates a bitmap of the QR Code
	/// // Requires a reference to System.Drawing
	/// using(var bitmap=QRCodeGenerator.EncodeToBitmap("http://www.example.com/")){
	///   // Saves the QR Code image to a file
	///   bitmap.Save("qrcode4.png");
	/// }
	/// </code>
	/// <para>If you define QRCODEPNG and the Png4BitIndexed class at
	///  http://articles.upokecenter.com/2011/05/11/png-4-bit-and-8-bit-indexed-encoder-classes-in-c/
	/// the following example is possible:</para>
	/// <code>
	/// // Generates a PNG object of the QR Code
	/// var png=QRCodeGenerator.Encode("http://www.example.com/");
	/// // Saves the QR Code image to a file
	/// png.Save("qrcode.png");
	/// </code>
	/// (QR Code is a trademark of Denso Wave Incorporated and may be registered
	/// in certain countries.)
	/// </summary>
	public static class QRCodeGenerator
	{
		
		private static int[] GaloisExp={
			1, 2, 4, 8, 16, 32, 64, 128, 29, 58, 116, 232, 205, 135, 19, 38, 76, 152, 45, 90, 180, 117, 234, 201, 143, 3, 6, 12, 24, 48, 96,
			192, 157, 39, 78, 156, 37, 74, 148, 53, 106, 212, 181, 119, 238, 193, 159, 35, 70, 140, 5, 10, 20, 40, 80, 160, 93, 186, 105, 210,
			185, 111, 222, 161, 95, 190, 97, 194, 153, 47, 94, 188, 101, 202, 137, 15, 30, 60, 120, 240, 253, 231, 211, 187, 107, 214, 177, 127,
			254, 225, 223, 163, 91, 182, 113, 226, 217, 175, 67, 134, 17, 34, 68, 136, 13, 26, 52, 104, 208, 189, 103, 206, 129, 31, 62, 124, 248,
			237, 199, 147, 59, 118, 236, 197, 151, 51, 102, 204, 133, 23, 46, 92, 184, 109, 218, 169, 79, 158, 33, 66, 132, 21, 42, 84, 168, 77, 154, 41, 82, 164, 85, 170, 73, 146, 57, 114, 228, 213, 183, 115, 230, 209, 191, 99, 198, 145, 63, 126, 252, 229, 215, 179, 123, 246, 241, 255, 227, 219, 171, 75, 150, 49, 98, 196, 149, 55, 110, 220, 165, 87, 174, 65, 130, 25, 50, 100, 200, 141, 7, 14, 28, 56, 112, 224, 221, 167, 83, 166, 81, 162, 89, 178, 121, 242, 249, 239, 195, 155, 43, 86, 172, 69, 138, 9, 18, 36, 72, 144, 61, 122, 244, 245, 247, 243, 251, 235, 203, 139, 11, 22, 44, 88, 176, 125, 250, 233, 207, 131, 27, 54, 108, 216, 173, 71, 142, 1, 2, 4, 8, 16, 32, 64, 128, 29, 58, 116, 232, 205, 135, 19, 38, 76, 152, 45, 90, 180, 117, 234, 201, 143, 3, 6, 12, 24, 48, 96, 192, 157, 39, 78, 156, 37, 74, 148, 53, 106, 212, 181, 119, 238, 193, 159, 35, 70, 140, 5, 10, 20, 40, 80, 160, 93, 186, 105, 210, 185, 111, 222, 161, 95, 190, 97, 194, 153, 47, 94, 188, 101, 202, 137, 15, 30, 60, 120, 240, 253, 231, 211, 187, 107, 214, 177, 127, 254, 225, 223, 163, 91, 182, 113, 226, 217, 175, 67, 134, 17, 34, 68, 136, 13, 26, 52, 104, 208, 189, 103, 206, 129, 31, 62, 124, 248, 237, 199, 147, 59, 118, 236, 197, 151, 51, 102, 204, 133, 23, 46, 92, 184, 109, 218, 169, 79, 158, 33, 66, 132, 21, 42, 84, 168, 77, 154, 41, 82, 164, 85, 170, 73, 146, 57, 114, 228, 213, 183, 115, 230, 209, 191, 99, 198, 145, 63, 126, 252, 229, 215, 179, 123, 246, 241, 255, 227, 219, 171, 75, 150, 49, 98, 196, 149, 55, 110, 220, 165, 87, 174, 65, 130, 25, 50, 100, 200, 141, 7, 14, 28, 56, 112, 224, 221, 167, 83, 166, 81, 162, 89, 178, 121, 242, 249, 239, 195, 155, 43, 86, 172, 69, 138, 9, 18, 36, 72, 144, 61, 122, 244, 245, 247, 243, 251, 235, 203, 139, 11, 22, 44, 88, 176, 125, 250, 233, 207, 131, 27, 54, 108, 216, 173, 71, 142, 1
		};
		private static int[] GaloisLog={
			0, 0, 1, 25, 2, 50, 26, 198, 3, 223, 51, 238, 27, 104, 199, 75, 4, 100, 224, 14, 52, 141, 239, 129, 28, 193, 105, 248, 200, 8, 76, 113, 5, 138, 101, 47, 225, 36, 15, 33, 53, 147, 142, 218, 240, 18, 130, 69, 29, 181, 194, 125, 106, 39, 249, 185, 201, 154, 9, 120, 77, 228, 114, 166, 6, 191, 139, 98, 102, 221, 48, 253, 226, 152, 37, 179, 16, 145, 34, 136, 54, 208, 148, 206, 143, 150, 219, 189, 241, 210, 19, 92, 131, 56, 70, 64, 30, 66, 182, 163, 195, 72, 126, 110, 107, 58, 40, 84, 250, 133, 186, 61, 202, 94, 155, 159, 10, 21, 121, 43, 78, 212, 229, 172, 115, 243, 167, 87, 7, 112, 192, 247, 140, 128, 99, 13, 103, 74, 222, 237, 49, 197, 254, 24, 227, 165, 153, 119, 38, 184, 180, 124, 17, 68, 146, 217, 35, 32, 137, 46, 55, 63, 209, 91, 149, 188, 207, 205, 144, 135, 151, 178, 220, 252, 190, 97, 242, 86, 211, 171, 20, 42, 93, 158, 132, 60, 57, 83, 71, 109, 65, 162, 31, 45, 67, 216, 183, 123, 164, 118, 196, 23, 73, 236, 127, 12, 111, 246, 108, 161, 59, 82, 41, 157, 85, 170, 251, 96, 134, 177, 187, 204, 62, 90, 203, 89, 95, 176, 156, 169, 160, 81, 11, 245, 22, 235, 122, 117, 44, 215, 79, 174, 213, 233, 230, 231, 173, 232, 116, 214, 244, 234, 168, 80, 88, 175
		};
		
		private static int[][] polynomials={
			null,
			null,
			null,
			null,
			null,
			null,
			null,
			new int[]{127, 122, 154, 164, 11, 68, 117},
			new int[]{255, 11, 81, 54, 239, 173, 200, 24},
			new int[]{226, 207, 158, 245, 235, 164, 232, 197, 37},
			new int[]{216, 194, 159, 111, 199, 94, 95, 113, 157, 193},
			new int[]{172, 130, 163, 50, 123, 219, 162, 248, 144, 116, 160},
			new int[]{68, 119, 67, 118, 220, 31, 7, 84, 92, 127, 213, 97},
			new int[]{137, 73, 227, 17, 177, 17, 52, 13, 46, 43, 83, 132, 120},
			new int[]{14, 54, 114, 70, 174, 151, 43, 158, 195, 127, 166, 210, 234, 163},
			new int[]{29, 196, 111, 163, 112, 74, 10, 105, 105, 139, 132, 151, 32, 134, 26},
			new int[]{59, 13, 104, 189, 68, 209, 30, 8, 163, 65, 41, 229, 98, 50, 36, 59},
			new int[]{119, 66, 83, 120, 119, 22, 197, 83, 249, 41, 143, 134, 85, 53, 125, 99, 79},
			new int[]{239, 251, 183, 113, 149, 175, 199, 215, 240, 220, 73, 82, 173, 75, 32, 67, 217, 146},
			null,
			new int[]{152, 185, 240, 5, 111, 99, 6, 220, 112, 150, 69, 36, 187, 22, 228, 198, 121, 121, 165, 174},
			null,
			new int[]{89, 179, 131, 176, 182, 244, 19, 189, 69, 40, 28, 137, 29, 123, 67, 253, 86, 218, 230, 26, 145, 245},
			null,
			new int[]{122, 118, 169, 70, 178, 237, 216, 102, 115, 150, 229, 73, 130, 72, 61, 43, 206, 1, 237, 247, 127, 217, 144, 117},
			null,
			new int[]{246, 51, 183, 4, 136, 98, 199, 152, 77, 56, 206, 24, 145, 40, 209, 117, 233, 42, 135, 68, 70, 144, 146, 77, 43, 94},
			null,
			new int[]{252, 9, 28, 13, 18, 251, 208, 150, 103, 174, 100, 41, 167, 12, 247, 56, 117, 119, 233, 127, 181, 100, 121, 147, 176, 74, 58, 197},
			null,
			new int[]{212, 246, 77, 73, 195, 192, 75, 98, 5, 70, 103, 177, 22, 217, 138, 51, 181, 246, 72, 25, 18, 46, 228, 74, 216, 195, 11, 106, 130, 150}
		};
		
		private static int GaloisMult(int a, int b){
			return (a==0 || b == 0) ? 0 : GaloisExp[GaloisLog[a]+GaloisLog[b]];
		}
		
		private static void ReedSolomonEncode(IList<byte> b, int offset, int count, int blockSize){
			int[] sr=new int[blockSize+1];
			int[] poly=polynomials[blockSize];
			if(poly==null)throw new InvalidOperationException();
			for (int i = 0; i < count; i++) {
				int dbyte = b[i+offset] ^ sr[0];
				for (int j = 0; j <blockSize-1; j++) {
					sr[j] = sr[j+1] ^ GaloisMult(poly[j], dbyte);
				}
				sr[blockSize-1] = GaloisMult(poly[blockSize-1], dbyte);
			}
			for(int i=0;i<blockSize;i++){
				b.Add((byte)sr[i]);
			}
		}
		
		private class CodeMatrix {
			byte[] data;
			int size;
			int codewordX;
			int codewordY;
			bool directionDown;
			bool directionRight;
			public CodeMatrix(int size){
				data=new byte[size*size];
				this.size=size;
				this.codewordX=size-1;
				this.codewordY=size-1;
			}
			public int Size {
				get { return size; }
			}
			public bool this[int x, int y]{
				get {
					return (data[y*size+x]&1)!=0;
				}
				set {
					data[y*size+x]=(value ? (byte)1 : (byte)0);
				}
			}
			
			private byte GetValue(int x, int y){
				return data[y*size+x];
			}
			
			private void SetValue(int x, int y, byte value){
				data[y*size+x]=value;
			}
			private void SetValueChecked(int x, int y, byte value){
				if(x>=0 && y>=0 && x<size && y<size){
					data[y*size+x]=value;
				}
			}
			
			public void SetVersionInfo(int value){
				for(int y=0;y<6;y++){
					for(int x=0;x<3;x++){
						SetValue(size-11+x,y,(byte)(0x50|(value&1)));
						SetValue(y,size-11+x,(byte)(0x50|(value&1)));
						value>>=1;
					}
				}
			}
			
			public void SetFormatInfo(int value){
				int value2=value;
				if(value==0){
					SetValue(8,size-8,0x40);
				} else {
					SetValue(8,size-8,0x41);
				}
				for(int y=0;y<6;y++){
					SetValue(8,y,(byte)(0x40|(value&1)));
					value>>=1;
				}
				SetValue(8,7,(byte)(0x40|(value&1)));
				SetValue(8,8,(byte)(0x40|((value>>1)&1)));
				SetValue(7,8,(byte)(0x40|((value>>2)&1)));
				value>>=3;
				for(int x=5;x>=0;x--){
					SetValue(x,8,(byte)(0x40|(value&1)));
					value>>=1;
				}
				for(int x=0;x<8;x++){
					SetValue(size-1-x,8,(byte)(0x40|(value2&1)));
					value2>>=1;
				}
				for(int y=0;y<7;y++){
					SetValue(8,size-7+y,(byte)(0x40|(value2&1)));
					value2>>=1;
				}
			}
			
			public void SetNextCodeword(byte value){
				for(int shift=0;shift<8;shift++){
					//Console.WriteLine("Drawing {0},{1} d={2} r={3}",codewordX,codewordY,
					//                  directionDown,directionRight);
					this[codewordX,codewordY]=(((value>>(7-shift))&1)!=0);
					int nextValue=0;
					do {
						if(directionDown && directionRight){
							codewordY++;
							if(codewordY>=size){
								codewordX--;
								if(codewordX==6)codewordX--;
								codewordY=size-1;
								directionDown=!directionDown;
							} else {
								codewordX++;
							}
							directionRight=!directionRight;
						} else if(directionDown){
							codewordX--;
							if(codewordX==6)codewordX--;
							directionRight=!directionRight;
						} else if(!directionDown && directionRight){
							codewordY--;
							if(codewordY<0){
								codewordX--;
								if(codewordX==6)codewordX--;
								codewordY=0;
								directionDown=!directionDown;
							} else {
								codewordX++;
							}
							directionRight=!directionRight;
						} else {
							codewordX--;
							if(codewordX==6)codewordX--;
							directionRight=!directionRight;
						}
						//Console.WriteLine("Checking {0},{1} d={2} r={3}",codewordX,codewordY,
						//                  directionDown,directionRight);
						if(codewordX<0)return;
						nextValue=GetValue(codewordX,codewordY);
					} while((nextValue>>1)!=0);
				}
			}

			public void DrawAlignmentPattern(int x, int y){
				if(x==0 || y==0 ||
				   (x<8 && y<8) ||
				   (x<8 && y>=size-8) ||
				   (y<8 && x>=size-8))
					return;
				x-=2;
				y-=2;
				for(int xx=x;xx<x+5;xx++){
					this.SetValue(xx,y,0x21);
					this.SetValue(xx,y+1,0x20);
					this.SetValue(xx,y+2,0x20);
					this.SetValue(xx,y+3,0x20);
					this.SetValue(xx,y+4,0x21);
				}
				for(int yy=y+1;yy<y+4;yy++){
					this.SetValue(x,yy,0x21);
					this.SetValue(x+4,yy,0x21);
				}
				this.SetValue(x+2,y+2,0x21);
			}
			
			private bool[] state11311expected={true,false,true,
				true,true,false,true,false};
			
			private int FindBlock(int x, int y, int tentativeWidth, byte[] moduleblocks){
				int mbs=x;
				int penalty=0;
				bool lastValue=this[x,y];
				//Console.WriteLine("tentative width: {0} x={1} y={2}",tentativeWidth,x,y);
				while(mbs<x+tentativeWidth){
					int blockHeight=1;
					int blockWidth=0;
					int mbsStep=0;
					for(int x2=mbs;x2<x+tentativeWidth;x2++){
						mbsStep++;
						if(this[x2,y+1]!=lastValue){
							break;
						}
						blockWidth++;
					}
					if(blockWidth>=2){
						blockHeight++;
						// Found a real block, now find its height
						int y2=0;
						for(y2=y+2;y2<size;y2++){
							bool haveBlock=true;
							for(int x2=mbs;x2<mbs+blockWidth;x2++){
								if(this[x2,y2]!=lastValue){
									haveBlock=false;
									break;
								}
							}
							if(haveBlock)
								blockHeight++;
							else
								break;
						}
						// mark the modules of the block found
						//Console.WriteLine("Found block: x={0} y={1} width={2} height={3}",mbs,y,blockWidth,blockHeight);
						int ypos=1;
						for(y2=y+1;y2<size && ypos<blockHeight;y2++,ypos++){
							for(int x2=mbs;x2<mbs+blockWidth;x2++){
								moduleblocks[y2*size+x2]=1;
							}
						}
						// add penalty
						penalty+=3*(blockHeight-1)*(blockWidth-1);
					}
					mbs+=mbsStep;
				}
				return penalty;
			}
			
			public int CalculatePenalty(){
				int dark=0;
				int adjacentTotal=0;
				bool lastValue=false;
				int state11311=0;
				int penalty=0;
				// Check rows and count dark modules
				for(int y=0;y<size;y++){
					adjacentTotal=0;
					lastValue=false;
					state11311=0;
					for(int x=0;x<size;x++){
						bool v=this[x,y];
						if(v){
							dark++;
						}
						if(v==state11311expected[state11311]){
							state11311++;
							if(state11311==8){
								state11311=0;
								//Console.WriteLine("Found 11311 row at x={0}, y={1}",x,y);
								penalty+=40;
							}
						} else {
							state11311=0;
						}
						if(adjacentTotal==0){
							adjacentTotal++;
							lastValue=v;
						} else if(lastValue==v){
							adjacentTotal++;
						} else {
							if(adjacentTotal>5){
								//Console.WriteLine("Found adjacent row: x={0}, y={1}, total={2}",x,y,adjacentTotal);
								penalty+=3*(adjacentTotal-5);
							}
							adjacentTotal=1;
							lastValue=v;
						}
					}
					// 1:1:3:1:1 pattern at end of row
					if(state11311==7){
						//Console.WriteLine("Found 11311 row at end, y={0}",y);
						penalty+=40;
					}
					if(adjacentTotal>5){
						//Console.WriteLine("Found adjacent row at end, y={0}, total={1}",y,adjacentTotal);
						penalty+=3*(adjacentTotal-5);
					}
				}
				// Check columns
				for(int x=0;x<size;x++){
					adjacentTotal=0;
					lastValue=false;
					state11311=0;
					for(int y=0;y<size;y++){
						bool v=this[x,y];
						if(v==state11311expected[state11311]){
							state11311++;
							if(state11311==8){
								//Console.WriteLine("Found 11311 column at x={0}, y={1}",x,y);
								state11311=0;
								penalty+=40;
							}
						} else {
							state11311=0;
						}
						if(adjacentTotal==0){
							adjacentTotal++;
							lastValue=v;
						} else if(lastValue==v){
							adjacentTotal++;
						} else {
							if(adjacentTotal>5){
								//Console.WriteLine("Found adjacent column: x={0}, y={1}, total={2}",x,y,adjacentTotal);
								penalty+=3*(adjacentTotal-5);
							}
							adjacentTotal=1;
							lastValue=v;
						}
					}
					// 1:1:3:1:1 pattern at end of row
					if(state11311==7){
						//Console.WriteLine("Found 11311 column at end, x={0}",x);
						penalty+=40;
					}
					if(adjacentTotal>5){
						//Console.WriteLine("Found adjacent column at end, x={0}, total={1}",x,adjacentTotal);
						penalty+=3*(adjacentTotal-5);
					}
				}
				// Count module blocks
				byte[] moduleblocks=new byte[size*size];
				int mbStart=0;
				for(int y=0;y<size-1;y++){
					adjacentTotal=0;
					lastValue=false;
					state11311=0;
					for(int x=0;x<size;x++){
						bool v=this[x,y];
						int mboffset=y*size+x;
						byte mb=moduleblocks[mboffset];
						if(adjacentTotal==0){
							if(mb==0){
								mbStart=x;
								adjacentTotal++;
								lastValue=v;
							}
						} else if(lastValue==v && mb==0){
							adjacentTotal++;
						} else {
							if(adjacentTotal>=2){
								// Found potential module block
								//Console.WriteLine("potential block at x={0},y={1}, v={2} lv={3}",x,y,v,lastValue);
								penalty+=FindBlock(mbStart,y,adjacentTotal,moduleblocks);
							}
							if(mb==0){
								mbStart=x;
								adjacentTotal=1;
								lastValue=v;
							} else {
								adjacentTotal=0;
							}
						}
						moduleblocks[mboffset]=1;
					}
					// Found potential module block at end of row
					if(adjacentTotal>=2){
						penalty+=FindBlock(mbStart,y,adjacentTotal,moduleblocks);
					}
				}
				// Proportion of dark modules
				int percent=100*dark/(size*size);
				percent=Math.Abs(percent-50)/5;
				penalty+=10*percent;
				//Console.WriteLine("Total penalty: {0}",penalty);
				return penalty;
			}
			
			public void ApplyMask(int mask){
				switch(mask){
					case 0:
						for(int y=0;y<size;y++){
							for(int x=0;x<size;x++){
								byte v=this.GetValue(x,y);
								if((v>>1)==0 && ((x+y)&1)==0){
									this.SetValue(x,y,(byte)(v^1));
								}
							}
						}
						break;
					case 1:
						for(int y=0;y<size;y+=2){
							for(int x=0;x<size;x++){
								byte v=this.GetValue(x,y);
								if((v>>1)==0){
									this.SetValue(x,y,(byte)(v^1));
								}
							}
						}
						break;
					case 2:
						for(int y=0;y<size;y++){
							for(int x=0;x<size;x+=3){
								byte v=this.GetValue(x,y);
								if((v>>1)==0){
									this.SetValue(x,y,(byte)(v^1));
								}
							}
						}
						break;
					case 3:
						for(int y=0;y<size;y++){
							for(int x=0;x<size;x++){
								byte v=this.GetValue(x,y);
								if((v>>1)==0 && ((x+y)%3)==0){
									this.SetValue(x,y,(byte)(v^1));
								}
							}
						}
						break;
					case 4:
						for(int y=0;y<size;y++){
							for(int x=0;x<size;x++){
								byte v=this.GetValue(x,y);
								if((v>>1)==0 && (((x/3)+(y>>1))&1)==0){
									this.SetValue(x,y,(byte)(v^1));
								}
							}
						}
						break;
					case 5:
						for(int y=0;y<size;y++){
							for(int x=0;x<size;x++){
								byte v=this.GetValue(x,y);
								int xy=x*y;
								if((v>>1)==0 && (((xy&1)+(xy%3)))==0){
									this.SetValue(x,y,(byte)(v^1));
								}
							}
						}
						break;
					case 6:
						for(int y=0;y<size;y++){
							for(int x=0;x<size;x++){
								byte v=this.GetValue(x,y);
								int xy=x*y;
								if((v>>1)==0 && (((xy&1)+(xy%3))&1)==0){
									this.SetValue(x,y,(byte)(v^1));
								}
							}
						}
						break;
					case 7:
						for(int y=0;y<size;y++){
							for(int x=0;x<size;x++){
								byte v=this.GetValue(x,y);
								int xy=x*y;
								if((v>>1)==0 && ((((x+y)&1)+(xy%3))&1)==0){
									this.SetValue(x,y,(byte)(v^1));
								}
							}
						}
						break;
					default:
						throw new ArgumentOutOfRangeException("mask");
				}
			}
			
			public void DrawTimingPatterns(){
				// Draw timing patterns
				for(int pos=0;pos<this.size-16;pos++){
					this.SetValue(pos+8,6,((byte)(((pos+1)&1)|0x30)));
					this.SetValue(6,pos+8,((byte)(((pos+1)&1)|0x30)));
				}
			}
			
			public void DrawFinderPattern(int x, int y){
				for(int yy=y-1;yy<y+8;yy++){
					for(int xx=x-1;xx<x+8;xx++){
						SetValueChecked(xx,yy,0x10);
					}
				}
				for(int xx=x;xx<x+7;xx++){
					SetValue(xx,y,0x11);
					SetValue(xx,y+6,0x11);
				}
				for(int yy=y+1;yy<y+6;yy++){
					SetValue(x,yy,0x11);
					SetValue(x+6,yy,0x11);
				}
				SetValue(x+2,y+2,0x11);
				SetValue(x+3,y+2,0x11);
				SetValue(x+4,y+2,0x11);
				SetValue(x+2,y+3,0x11);
				SetValue(x+3,y+3,0x11);
				SetValue(x+4,y+3,0x11);
				SetValue(x+2,y+4,0x11);
				SetValue(x+3,y+4,0x11);
				SetValue(x+4,y+4,0x11);
			}
			
			public override string ToString()
			{
				System.Text.StringBuilder sb=new System.Text.StringBuilder();
				for(int y=0;y<this.size;y++){
					for(int x=0;x<this.size;x++){
						sb.Append(this[x,y] ? '*' : ' ');
					}
					sb.Append(Environment.NewLine);
				}
				return sb.ToString();
			}

			
		}
		
		private class BitList {
			List<byte> list=new List<byte>();
			byte curbyte;
			int curbit;
			public BitList(){}
			public IList<byte> List {
				get { return list; }
			}
			public void AddBits(byte value, int bits){
				for(int i=bits-1;i>=0;i--){
					curbyte|=(byte)(((value>>i)&(byte)1)<<(7-curbit));
					curbit++;
					if(curbit>=8){
						curbit=0;
						list.Add(curbyte);
						curbyte=0;
					}
				}
			}
			public void AddBits(int value, int bits){
				for(int i=bits-1;i>=0;i--){
					curbyte|=(byte)(((value>>i)&(byte)1)<<(7-curbit));
					curbit++;
					if(curbit>=8){
						curbit=0;
						list.Add(curbyte);
						curbyte=0;
					}
				}
			}
			public void PadTo(int size){
				// Add terminator
				if(list.Count+1<size || curbit<=4){
					AddBits(0,4);
				}
				if(curbit>0){					
					list.Add(curbyte);
					curbit=0;
				}
				// Add padding bytes
				curbyte=0xEC;
				while(list.Count<size){
					list.Add(curbyte);
					curbyte=(curbyte==0xEC) ? (byte)0x11 : (byte)0xEC;
				}
				curbyte=0;
			}
		}
		
		static int[] VersionInfo={
			0x7c94,0x85bc,0x9a99,0xa4d3,
			0xbbf6,0xc762,0xd847,0xe60d,0xf928,
			0x10b78,0x1145d,0x12a17,0x13532,0x149a6,
			0x15683,0x168c9,0x177ec,0x18ec4,0x191e1,
			0x1afab,0x1b08e,0x1cc1a,0x1d33f,0x1ed75,
			0x1f250,0x209d5,0x216f0,0x228ba,0x2379f,
			0x24b0b,0x2542e,0x26a64,0x27541,0x28c69
		};
		
		static int[] AlignmentPatterns={
			0,0,0,0,0,0,0,
			6,18,0,0,0,0,0,
			6,22,0,0,0,0,0,
			6,26,0,0,0,0,0,
			6,30,0,0,0,0,0,
			6,34,0,0,0,0,0,
			6,22,38,0,0,0,0,
			6,24,42,0,0,0,0,
			6,26,46,0,0,0,0,
			6,28,50,0,0,0,0,
			6,30,54,0,0,0,0,
			6,32,58,0,0,0,0,
			6,34,62,0,0,0,0,
			6,26,46,66,0,0,0,
			6,26,48,70,0,0,0,
			6,26,50,74,0,0,0,
			6,30,54,78,0,0,0,
			6,30,56,82,0,0,0,
			6,30,58,86,0,0,0,
			6,34,62,90,0,0,0,
			6,28,50,72,94,0,0,
			6,26,50,74,98,0,0,
			6,30,54,78,102,0,0,
			6,28,54,80,106,0,0,
			6,32,58,84,110,0,0,
			6,30,58,86,114,0,0,
			6,34,62,90,118,0,0,
			6,26,50,74,98,122,0,
			6,30,54,78,102,126,0,
			6,26,52,78,104,130,0,
			6,30,56,82,108,134,0,
			6,34,60,86,112,138,0,
			6,30,58,86,114,142,0,
			6,34,62,90,118,146,0,
			6,30,54,78,102,126,150,
			6,24,50,76,102,128,154,
			6,28,54,80,106,132,158,
			6,32,58,84,110,136,162,
			6,26,54,82,110,138,166,
			6,30,58,86,114,142,170
		};
		
		static int[] DataBlocks={
			1,16,10,0,0,0,
			1,28,16,0,0,0,
			1,44,26,0,0,0,
			2,32,18,0,0,0,
			2,43,24,0,0,0,
			4,27,16,0,0,0,
			4,31,18,0,0,0,
			2,38,22,2,39,22,
			3,36,22,2,37,22,
			4,43,26,1,44,26,
			1,50,30,4,51,30,
			6,36,22,2,37,22,
			8,37,22,1,38,22,
			4,40,24,5,41,24,
			5,41,24,5,42,24,
			7,45,28,3,46,28,
			10,46,28,1,47,28,
			9,43,26,4,44,26,
			3,44,26,11,45,26,
			3,41,26,13,42,26,
			17,42,26,0,0,0,
			17,46,28,0,0,0,
			4,47,28,14,48,28,
			6,45,28,14,46,28,
			8,47,28,13,48,28,
			19,46,28,4,47,28,
			22,45,28,3,46,28,
			3,45,28,23,46,28,
			21,45,28,7,46,28,
			19,47,28,10,48,28,
			2,46,28,29,47,28,
			10,46,28,23,47,28,
			14,46,28,21,47,28,
			14,46,28,23,47,28,
			12,47,28,26,48,28,
			6,47,28,34,48,28,
			29,46,28,14,47,28,
			13,46,28,32,47,28,
			40,47,28,7,48,28,
			18,47,28,31,48,28,
		};
		
		static int[] FormatInfo=new int[]{
			21522, 20773, 24188, 23371, 17913, 16590,
			20375, 19104, 30660, 29427, 32170, 30877,
			26159, 25368, 27713, 26998, 5769, 5054,
			7399, 6608, 1890, 597, 3340, 2107, 13663,
			12392, 16177, 14854, 9396, 8579, 11994, 11245
		};
		
		static int[] versions=new int[]{
			14,26,42,62,84,106,
			122,152,180,213,251,287,
			331,362,412,450,504,560,
			624,666,711,779,857,911,
			997,1059,1125,1190,1264,1370,
			1452,1538,1628,1722,1809,1911,
			1989,2099,2213,2331
		};
		
		static int[] versionsnumeric={
			34,63,101,149,
			202,255,293,365,432,513,
			604,691,796,871,991,1082,
			1212,1346,1500,1600,1708,1872,
			2059,2188,2395,2544,2701,2857,
			3035,3289,3486,3693,3909,4134,
			4343,4588,4775,5039,5313,5596
		};
		
		static int[] versionsalphanum={
			20,38,61,90,122,154,
			178,221,262,311,366,419,
			483,528,600,656,734,816,
			909,970,1035,1134,1248,1326,
			1451,1542,1637,1732,1839,1994,
			2113,2238,2369,2506,2632,2780,
			2894,3054,3220,3391
		};
		
		static int[] ecbytes=new int[]{
			10,16,26,36,48,64,
			72,88,110,130,150,176,
			198,216,240,280,308,338,
			364,416,442,476,504,560,
			588,644,700,728,784,812,
			868,924,980,1036,1064,1120,
			1204,1260,1316,1372
		};
		
		static int[] CharToValue={
			36,0,0,0,37,38,0,0,0,0,39,40,0,41,42,43,
			0,1,2,3,4,5,6,7,8,9,44,0,0,0,0,0,
			0,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,
			25,26,27,28,29,30,31,32,33,34,35,0,0,0,0,0
		};
		
		private static CodeMatrix GetEncodedMatrix(byte[] data, bool stringData){
			if(data==null)throw new ArgumentNullException("data");
			if(data.Length>7089)throw new ArgumentException("'data' length is too long.");
			bool allnumeric=stringData;
			bool alphanumeric=stringData;
			int version=-1;
			if(stringData){
				for(int i=0;i<data.Length;i++){
					byte d=data[i];
					if(d<0x30 || d>=0x3a){
						allnumeric=false;
					}
					if(d<0x20 || d>=0x60 || CharToValue[d-0x20]==0){
						alphanumeric=false;
					}
				}
			}
			//Console.WriteLine("{0} {1}",alphanumeric,stringData);
			if(allnumeric){
				for(int i=0;i<versionsnumeric.Length;i++){
					if(data.Length<=versionsnumeric[i]){
						version=i;
						break;
					}
				}
			} else if(alphanumeric){
				for(int i=0;i<versionsalphanum.Length;i++){
					if(data.Length<=versionsalphanum[i]){
						version=i;
						break;
					}
				}
			} else {
				for(int i=0;i<versions.Length;i++){
					if(data.Length<=versions[i]){
						version=i;
						break;
					}
				}
			}
			if(version<0){
				throw new ArgumentException("'data' length is too long.");
			}
			BitList bl=new BitList();
			if(allnumeric){
				bl.AddBits((byte)1,4);
				int datacount=10;
				if(version>=9 && version<=25)datacount=12;
				if(version>=26)datacount=14;
				bl.AddBits(data.Length,datacount);
				int runningcount=0;
				int runningdigits=0;
				for(int i=0;i<data.Length;i++){
					runningdigits*=10;
					runningdigits+=((int)data[i]-0x30);
					runningcount++;
					if(runningcount==3){
						bl.AddBits(runningdigits,10);
						runningcount=0;
						runningdigits=0;
					}
				}
				if(runningcount>0){
					bl.AddBits(runningdigits,(runningcount==1) ? 4 : 7);
				}
			} else if(alphanumeric){
				bl.AddBits((byte)2,4);
				int datacount=9;
				if(version>=9 && version<=25)datacount=11;
				if(version>=26)datacount=13;
				bl.AddBits(data.Length,datacount);
				//Console.WriteLine("{0} ver={1}",data.Length,version+1);
				int runningcount=0;
				int runningdigits=0;
				for(int i=0;i<data.Length;i++){
					runningdigits*=45;
					runningdigits+=CharToValue[data[i]-0x20];
					runningcount++;
					if(runningcount==2){
						bl.AddBits(runningdigits,11);
						runningcount=0;
						runningdigits=0;
					}
				}
				if(runningcount>0){
					bl.AddBits(runningdigits,6);
				}
			} else {
				bl.AddBits((byte)4,4);
				bl.AddBits(data.Length,(version<9) ? 8 : 16);
				for(int i=0;i<data.Length;i++){
					bl.AddBits(data[i],8);
				}
			}
			int matrixSize=21+version*4;
			int dataWords=(version<9) ? versions[version]+2 : versions[version]+3;
			bl.PadTo(dataWords);
			//Console.WriteLine(ArrayUtil.ArrayToString(bl.List));
			int rsOffset=version*6;
			int dataOffset=0;
			int ecOffset=0;
			int firstBlockCount=DataBlocks[rsOffset];
			int secondBlockCount=DataBlocks[rsOffset+3];
			int totalBlockCount=firstBlockCount+secondBlockCount;
			// first all data blocks then all error correction blocks
			int[] codewordOffsets=new int[totalBlockCount*2];
			for(int i=0;i<firstBlockCount;i++){
				ReedSolomonEncode(bl.List,dataOffset,
				                  DataBlocks[rsOffset+1],
				                  DataBlocks[rsOffset+2]);
				codewordOffsets[i]=dataOffset;
				codewordOffsets[totalBlockCount+i]=dataWords+ecOffset;
				dataOffset+=DataBlocks[rsOffset+1];
				ecOffset+=DataBlocks[rsOffset+2];
			}
			for(int i=0;i<secondBlockCount;i++){
				ReedSolomonEncode(bl.List,dataOffset,
				                  DataBlocks[rsOffset+4],
				                  DataBlocks[rsOffset+5]);
				codewordOffsets[i+DataBlocks[rsOffset]]=dataOffset;
				codewordOffsets[i+totalBlockCount+DataBlocks[rsOffset]]=dataWords+ecOffset;
				dataOffset+=DataBlocks[rsOffset+4];
			}
			//Console.WriteLine(ArrayUtil.ArrayToString(bl.List));
			CodeMatrix matrix=new CodeMatrix(matrixSize);
			matrix.DrawFinderPattern(0,0);
			matrix.DrawFinderPattern(0,matrixSize-7);
			matrix.DrawFinderPattern(matrixSize-7,0);
			matrix.DrawTimingPatterns();
			matrix.SetFormatInfo(0);
			if(version>=6){
				matrix.SetVersionInfo(0);
			}
			int alignOffset=version*7;
			for(int y=0;y<7;y++){
				for(int x=0;x<7;x++){
					matrix.DrawAlignmentPattern(AlignmentPatterns[alignOffset+x],
					                            AlignmentPatterns[alignOffset+y]);
				}
			}
			// Set data words
			int minSizePerBlock=DataBlocks[rsOffset+1];
			for(int i=0;i<minSizePerBlock;i++){
				for(int j=0;j<totalBlockCount;j++){
					//Console.WriteLine("{0}->{0:X2}",codewordOffsets[j],bl.List[codewordOffsets[j]]);
					int index=codewordOffsets[j];
					matrix.SetNextCodeword(bl.List[index]);
					codewordOffsets[j]++;
				}
			}
			if(secondBlockCount!=0){
				for(int j=0;j<secondBlockCount;j++){
					int index=codewordOffsets[j+firstBlockCount];
					matrix.SetNextCodeword(bl.List[index]);
					codewordOffsets[j+firstBlockCount]++;
				}
			}
			// Set error correction words
			minSizePerBlock=DataBlocks[rsOffset+2];
			for(int i=0;i<minSizePerBlock;i++){
				for(int j=0;j<totalBlockCount;j++){
					int index=codewordOffsets[j+totalBlockCount];
					matrix.SetNextCodeword(bl.List[index]);
					codewordOffsets[j+totalBlockCount]++;
				}
			}
			// Mask the data
			int lowestPenalty=0;
			int bestMask=0;
			for(int i=0;i<8;i++){
				matrix.ApplyMask(i);
				int penalty=matrix.CalculatePenalty();
				if(i==0 || penalty<lowestPenalty){
					bestMask=i;
					lowestPenalty=penalty;
				}
				// Reapply the mask to erase it
				matrix.ApplyMask(i);
			}
			matrix.ApplyMask(bestMask);
			matrix.SetFormatInfo(FormatInfo[bestMask]);
			//Console.WriteLine("format={0}",FormatInfo[bestMask]);
			if(version>=6){
				//	Console.WriteLine("version={0}",version+1);
				//	Console.WriteLine("version={0:X5}",VersionInfo[version-6]);
				matrix.SetVersionInfo(VersionInfo[version-6]);
			}
			return matrix;
		}
		
		#if QRCODESYSTEMDRAWING
		private static System.Drawing.Bitmap EncodeToBitmap(CodeMatrix matrix){
			int moduleSize=6;
			int pngSize=moduleSize*(matrix.Size+8);
			while(moduleSize>2 && pngSize>=600){
				moduleSize--;
				pngSize=moduleSize*(matrix.Size+8);
			}
			System.Drawing.Bitmap image = new System.Drawing.Bitmap(pngSize,pngSize);
			using(System.Drawing.Graphics g = System.Drawing.Graphics.FromImage(image)){
				g.FillRectangle(System.Drawing.Brushes.White, new System.Drawing.Rectangle(0, 0, image.Width, image.Height));
				for(int y=0;y<matrix.Size;y++){
					for(int x=0;x<matrix.Size;x++){
						int pngx=(4+x)*moduleSize;
						int pngy=(4+y)*moduleSize;
						bool v=matrix[x,y];
						int color=(v ? 1 : 0);
						if(v){
							g.FillRectangle(System.Drawing.Brushes.Black,
							                new System.Drawing.Rectangle(pngx,pngy,moduleSize,moduleSize));
						}
					}
				}
			}
			return image;
		}

		/// <summary>
		/// Generates a QR Code bitmap with the data set to the specified string, in UTF-8 encoding.
		/// </summary>
		/// <param name="data">A string to encode as a QR Code</param>
		/// <returns>A bitmap whose image is a QR Code.</returns>
		/// <exception cref="T:System.ArgumentNullException">'data' is null.</exception>
		/// <exception cref="T:System.ArgumentNullException">'data' is too long to fit in a QR Code.</exception>
		public static System.Drawing.Bitmap EncodeToBitmap(string data){
			if(data==null)throw new ArgumentNullException("data");
			return EncodeToBitmap(GetEncodedMatrix(System.Text.Encoding.UTF8.GetBytes(data),true));
		}
		
		/// <summary>
		/// Generates a QR Code bitmap with the data set to the specified array of bytes.
		/// </summary>
		/// <param name="data">A string to encode as a QR Code</param>
		/// <returns>A bitmap whose image is a QR Code.</returns>
		/// <exception cref="T:System.ArgumentNullException">'data' is null.</exception>
		/// <exception cref="T:System.ArgumentNullException">'data' is too long to fit in a QR Code.</exception>
		public static System.Drawing.Bitmap EncodeToBitmap(byte[] data){
			return EncodeToBitmap(GetEncodedMatrix(data,false));
		}
		#endif
		#if QRCODEPNG
		private static PeterO.Png4BitIndexed Encode(CodeMatrix matrix){
			int moduleSize=6;
			int pngSize=moduleSize*(matrix.Size+8);
			while(moduleSize>2 && pngSize>=600){
				moduleSize--;
				pngSize=moduleSize*(matrix.Size+8);
			}
			PeterO.Png4BitIndexed png=new Png4BitIndexed(pngSize,pngSize);
			png.SetColor(0,new byte[]{255,255,255,255});
			png.SetColor(1,new byte[]{0,0,0,255});
			for(int y=0;y<png.Height;y++){
				for(int x=0;x<png.Width;x++){
					png.SetPixel(x,y,0);
				}
			}
			for(int y=0;y<matrix.Size;y++){
				for(int x=0;x<matrix.Size;x++){
					int pngx=(4+x)*moduleSize;
					int pngy=(4+y)*moduleSize;
					bool v=matrix[x,y];
					int color=(v ? 1 : 0);
					for(int y2=0;y2<moduleSize;y2++){
						for(int x2=0;x2<moduleSize;x2++){
							png.SetPixel(pngx+x2,pngy+y2,color);
						}
					}
				}
			}
			return png;
		}
		
		/// <summary>
		/// Generates a QR Code image with the data set to the specified string, in UTF-8 encoding.
		/// </summary>
		/// <param name="data">A string to encode as a QR Code</param>
		/// <returns>A PNG object whose image is a QR Code.</returns>
		/// <exception cref="T:System.ArgumentNullException">'data' is null.</exception>
		/// <exception cref="T:System.ArgumentNullException">'data' is too long to fit in a QR Code.</exception>
		public static PeterO.Png4BitIndexed Encode(string data){
			if(data==null)throw new ArgumentNullException("data");
			return Encode(GetEncodedMatrix(System.Text.Encoding.UTF8.GetBytes(data),true));
		}

		/// <summary>
		/// Generates a QR Code image with the data set to the specified array of bytes.
		/// </summary>
		/// <param name="data">A string to encode as a QR Code</param>
		/// <returns>A PNG object whose image is a QR Code.</returns>
		/// <exception cref="T:System.ArgumentNullException">'data' is null.</exception>
		/// <exception cref="T:System.ArgumentNullException">'data' is too long to fit in a QR Code.</exception>
		public static PeterO.Png4BitIndexed Encode(byte[] data){
			if(data==null)throw new ArgumentNullException("data");
			return Encode(GetEncodedMatrix(data,false));
		}
		#endif
	}
}
