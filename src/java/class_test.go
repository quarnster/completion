package java

import (
	"archive/zip"
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
	"testing"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func rtJar(t *testing.T) string {
	substr := []byte("rt.jar")

	// -XshowSettings still exits with `1` returning an error, so use CombinedOutput.
	out, _ := exec.Command("java", "-XshowSettings:properties").CombinedOutput()

	for _, line := range bytes.Split(out, []byte("\n")) {
		if bytes.Contains(line, substr) {
			return string(bytes.TrimSpace(line))
		}
	}

	t.Fatal("Failed to locate rt.jar")
	panic("unreachable")
}

func TestLoadAllClasses(t *testing.T) {
	if z, err := zip.OpenReader(rtJar(t)); err != nil {
		t.Fatal(err)
	} else {
		defer z.Close()
		var (
			inChan  = make(chan []byte, len(z.File))
			outChan = make(chan error, 32)
			wg      sync.WaitGroup
			once    sync.Once
		)
		worker := func() {
			wg.Add(1)
			for zf := range inChan {
				f := bytes.NewReader(zf)
				if c, err := NewClass(f); err != nil {
					outChan <- err
				} else {
					t.Log("class", String(c.Constant_pool, c.This_class))
				}
			}
			wg.Done()
			wg.Wait()
			once.Do(func() {
				close(outChan)
			})
		}
		for i := 0; i < runtime.NumCPU(); i++ {
			go worker()
		}
		for _, zf := range z.File {
			if strings.HasSuffix(zf.Name, ".class") {
				if f, err := zf.Open(); err != nil {
					t.Error(err)
				} else {
					defer f.Close()
					if d, err := ioutil.ReadAll(f); err != nil {
						t.Error(err)
					} else {
						inChan <- d
					}
				}
			}
		}
		close(inChan)
		for o := range outChan {
			if o != nil {
				t.Error(o)
			}
		}
	}
}

// Diff cut'n'paste from http://golang.org/src/cmd/gofmt/gofmt.go
func diff(b1, b2 []byte) (data []byte, err error) {
	f1, err := ioutil.TempFile("", "parser")
	if err != nil {
		return
	}
	defer os.Remove(f1.Name())
	defer f1.Close()

	f2, err := ioutil.TempFile("", "parser")
	if err != nil {
		return
	}
	defer os.Remove(f2.Name())
	defer f2.Close()

	f1.Write(b1)
	f2.Write(b2)

	data, err = exec.Command("diff", "-u", f1.Name(), f2.Name()).CombinedOutput()
	if len(data) > 0 {
		// diff exits with a non-zero status when the files don't match.
		// Ignore that failure as long as we get output.
		err = nil
	}
	return

}

func TestSpecificClasses(t *testing.T) {
	tests := map[string][]byte{
		"java/lang/String.class": []byte(`public final class java/lang/String
extends java/lang/Object
implements
	java/io/Serializable
	java/lang/Comparable
	java/lang/CharSequence
Fields
	private final value [C
	private hash I
	private static final serialVersionUID J
	private static final serialPersistentFields [Ljava/io/ObjectStreamField;
	public static final CASE_INSENSITIVE_ORDER Ljava/util/Comparator;
	private static final HASHING_SEED I
	private transient hash32 I
Methods
	public <init> ()V
	public <init> (Ljava/lang/String;)V
	public <init> ([C)V
	public <init> ([CII)V
	public <init> ([III)V
	public <init> ([BIII)V
	public <init> ([BI)V
	private static checkBounds ([BII)V
	public <init> ([BIILjava/lang/String;)V
	public <init> ([BIILjava/nio/charset/Charset;)V
	public <init> ([BLjava/lang/String;)V
	public <init> ([BLjava/nio/charset/Charset;)V
	public <init> ([BII)V
	public <init> ([B)V
	public <init> (Ljava/lang/StringBuffer;)V
	public <init> (Ljava/lang/StringBuilder;)V
	<init> ([CZ)V
	<init> (II[C)V
	public length ()I
	public isEmpty ()Z
	public charAt (I)C
	public codePointAt (I)I
	public codePointBefore (I)I
	public codePointCount (II)I
	public offsetByCodePoints (II)I
	getChars ([CI)V
	public getChars (II[CI)V
	public getBytes (II[BI)V
	public getBytes (Ljava/lang/String;)[B
	public getBytes (Ljava/nio/charset/Charset;)[B
	public getBytes ()[B
	public equals (Ljava/lang/Object;)Z
	public contentEquals (Ljava/lang/StringBuffer;)Z
	public contentEquals (Ljava/lang/CharSequence;)Z
	public equalsIgnoreCase (Ljava/lang/String;)Z
	public compareTo (Ljava/lang/String;)I
	public compareToIgnoreCase (Ljava/lang/String;)I
	public regionMatches (ILjava/lang/String;II)Z
	public regionMatches (ZILjava/lang/String;II)Z
	public startsWith (Ljava/lang/String;I)Z
	public startsWith (Ljava/lang/String;)Z
	public endsWith (Ljava/lang/String;)Z
	public hashCode ()I
	public indexOf (I)I
	public indexOf (II)I
	private indexOfSupplementary (II)I
	public lastIndexOf (I)I
	public lastIndexOf (II)I
	private lastIndexOfSupplementary (II)I
	public indexOf (Ljava/lang/String;)I
	public indexOf (Ljava/lang/String;I)I
	static indexOf ([CII[CIII)I
	public lastIndexOf (Ljava/lang/String;)I
	public lastIndexOf (Ljava/lang/String;I)I
	static lastIndexOf ([CII[CIII)I
	public substring (I)Ljava/lang/String;
	public substring (II)Ljava/lang/String;
	public subSequence (II)Ljava/lang/CharSequence;
	public concat (Ljava/lang/String;)Ljava/lang/String;
	public replace (CC)Ljava/lang/String;
	public matches (Ljava/lang/String;)Z
	public contains (Ljava/lang/CharSequence;)Z
	public replaceFirst (Ljava/lang/String;Ljava/lang/String;)Ljava/lang/String;
	public replaceAll (Ljava/lang/String;Ljava/lang/String;)Ljava/lang/String;
	public replace (Ljava/lang/CharSequence;Ljava/lang/CharSequence;)Ljava/lang/String;
	public split (Ljava/lang/String;I)[Ljava/lang/String;
	public split (Ljava/lang/String;)[Ljava/lang/String;
	public toLowerCase (Ljava/util/Locale;)Ljava/lang/String;
	public toLowerCase ()Ljava/lang/String;
	public toUpperCase (Ljava/util/Locale;)Ljava/lang/String;
	public toUpperCase ()Ljava/lang/String;
	public trim ()Ljava/lang/String;
	public toString ()Ljava/lang/String;
	public toCharArray ()[C
	public static transient format (Ljava/lang/String;[Ljava/lang/Object;)Ljava/lang/String;
	public static transient format (Ljava/util/Locale;Ljava/lang/String;[Ljava/lang/Object;)Ljava/lang/String;
	public static valueOf (Ljava/lang/Object;)Ljava/lang/String;
	public static valueOf ([C)Ljava/lang/String;
	public static valueOf ([CII)Ljava/lang/String;
	public static copyValueOf ([CII)Ljava/lang/String;
	public static copyValueOf ([C)Ljava/lang/String;
	public static valueOf (Z)Ljava/lang/String;
	public static valueOf (C)Ljava/lang/String;
	public static valueOf (I)Ljava/lang/String;
	public static valueOf (J)Ljava/lang/String;
	public static valueOf (F)Ljava/lang/String;
	public static valueOf (D)Ljava/lang/String;
	public intern ()Ljava/lang/String;
	hash32 ()I
	public volatile synthetic compareTo (Ljava/lang/Object;)I
	static <clinit> ()V
`),
		"java/util/ArrayList.class": []byte(`public class java/util/ArrayList
extends java/util/AbstractList
implements
	java/util/List
	java/util/RandomAccess
	java/lang/Cloneable
	java/io/Serializable
Fields
	private static final serialVersionUID J
	private transient elementData [Ljava/lang/Object;
	private size I
	private static final MAX_ARRAY_SIZE I
Methods
	public <init> (I)V
	public <init> ()V
	public <init> (Ljava/util/Collection;)V
	public trimToSize ()V
	public ensureCapacity (I)V
	private ensureCapacityInternal (I)V
	private grow (I)V
	private static hugeCapacity (I)I
	public size ()I
	public isEmpty ()Z
	public contains (Ljava/lang/Object;)Z
	public indexOf (Ljava/lang/Object;)I
	public lastIndexOf (Ljava/lang/Object;)I
	public clone ()Ljava/lang/Object;
	public toArray ()[Ljava/lang/Object;
	public toArray ([Ljava/lang/Object;)[Ljava/lang/Object;
	elementData (I)Ljava/lang/Object;
	public get (I)Ljava/lang/Object;
	public set (ILjava/lang/Object;)Ljava/lang/Object;
	public add (Ljava/lang/Object;)Z
	public add (ILjava/lang/Object;)V
	public remove (I)Ljava/lang/Object;
	public remove (Ljava/lang/Object;)Z
	private fastRemove (I)V
	public clear ()V
	public addAll (Ljava/util/Collection;)Z
	public addAll (ILjava/util/Collection;)Z
	protected removeRange (II)V
	private rangeCheck (I)V
	private rangeCheckForAdd (I)V
	private outOfBoundsMsg (I)Ljava/lang/String;
	public removeAll (Ljava/util/Collection;)Z
	public retainAll (Ljava/util/Collection;)Z
	private batchRemove (Ljava/util/Collection;Z)Z
	private writeObject (Ljava/io/ObjectOutputStream;)V
	private readObject (Ljava/io/ObjectInputStream;)V
	public listIterator (I)Ljava/util/ListIterator;
	public listIterator ()Ljava/util/ListIterator;
	public iterator ()Ljava/util/Iterator;
	public subList (II)Ljava/util/List;
	static subListRangeCheck (III)V
	static synthetic access$100 (Ljava/util/ArrayList;)I
	static synthetic access$200 (Ljava/util/ArrayList;)[Ljava/lang/Object;
`),
		"java/awt/Button.class": []byte(`public class java/awt/Button
extends java/awt/Component
implements
	javax/accessibility/Accessible
Fields
	label Ljava/lang/String;
	actionCommand Ljava/lang/String;
	transient actionListener Ljava/awt/event/ActionListener;
	private static final base Ljava/lang/String;
	private static nameCounter I
	private static final serialVersionUID J
	private buttonSerializedDataVersion I
Methods
	private static initIDs ()V
	public <init> ()V
	public <init> (Ljava/lang/String;)V
	constructComponentName ()Ljava/lang/String;
	public addNotify ()V
	public getLabel ()Ljava/lang/String;
	public setLabel (Ljava/lang/String;)V
	public setActionCommand (Ljava/lang/String;)V
	public getActionCommand ()Ljava/lang/String;
	public addActionListener (Ljava/awt/event/ActionListener;)V
	public removeActionListener (Ljava/awt/event/ActionListener;)V
	public getActionListeners ()[Ljava/awt/event/ActionListener;
	public getListeners (Ljava/lang/Class;)[Ljava/util/EventListener;
	eventEnabled (Ljava/awt/AWTEvent;)Z
	protected processEvent (Ljava/awt/AWTEvent;)V
	protected processActionEvent (Ljava/awt/event/ActionEvent;)V
	protected paramString ()Ljava/lang/String;
	private writeObject (Ljava/io/ObjectOutputStream;)V
	private readObject (Ljava/io/ObjectInputStream;)V
	public getAccessibleContext ()Ljavax/accessibility/AccessibleContext;
	static <clinit> ()V
`),
		"javax/swing/JLabel.class": []byte(`public class javax/swing/JLabel
extends javax/swing/JComponent
implements
	javax/swing/SwingConstants
	javax/accessibility/Accessible
Fields
	private static final uiClassID Ljava/lang/String;
	private mnemonic I
	private mnemonicIndex I
	private text Ljava/lang/String;
	private defaultIcon Ljavax/swing/Icon;
	private disabledIcon Ljavax/swing/Icon;
	private disabledIconSet Z
	private verticalAlignment I
	private horizontalAlignment I
	private verticalTextPosition I
	private horizontalTextPosition I
	private iconTextGap I
	protected labelFor Ljava/awt/Component;
	static final LABELED_BY_PROPERTY Ljava/lang/String;
Methods
	public <init> (Ljava/lang/String;Ljavax/swing/Icon;I)V
	public <init> (Ljava/lang/String;I)V
	public <init> (Ljava/lang/String;)V
	public <init> (Ljavax/swing/Icon;I)V
	public <init> (Ljavax/swing/Icon;)V
	public <init> ()V
	public getUI ()Ljavax/swing/plaf/LabelUI;
	public setUI (Ljavax/swing/plaf/LabelUI;)V
	public updateUI ()V
	public getUIClassID ()Ljava/lang/String;
	public getText ()Ljava/lang/String;
	public setText (Ljava/lang/String;)V
	public getIcon ()Ljavax/swing/Icon;
	public setIcon (Ljavax/swing/Icon;)V
	public getDisabledIcon ()Ljavax/swing/Icon;
	public setDisabledIcon (Ljavax/swing/Icon;)V
	public setDisplayedMnemonic (I)V
	public setDisplayedMnemonic (C)V
	public getDisplayedMnemonic ()I
	public setDisplayedMnemonicIndex (I)V
	public getDisplayedMnemonicIndex ()I
	protected checkHorizontalKey (ILjava/lang/String;)I
	protected checkVerticalKey (ILjava/lang/String;)I
	public getIconTextGap ()I
	public setIconTextGap (I)V
	public getVerticalAlignment ()I
	public setVerticalAlignment (I)V
	public getHorizontalAlignment ()I
	public setHorizontalAlignment (I)V
	public getVerticalTextPosition ()I
	public setVerticalTextPosition (I)V
	public getHorizontalTextPosition ()I
	public setHorizontalTextPosition (I)V
	public imageUpdate (Ljava/awt/Image;IIIII)Z
	private writeObject (Ljava/io/ObjectOutputStream;)V
	protected paramString ()Ljava/lang/String;
	public getLabelFor ()Ljava/awt/Component;
	public setLabelFor (Ljava/awt/Component;)V
	public getAccessibleContext ()Ljavax/accessibility/AccessibleContext;
`),
	}

	var err error

	z, err := zip.OpenReader(rtJar(t))
	if err != nil {
		t.Fatalf("Failed to read rt.jar: %s", err)
	}
	defer z.Close()

	for _, zf := range z.File {

		v, ok := tests[zf.Name]
		if !ok {
			continue
		}

		f, err := zf.Open()
		if err != nil {
			t.Error("Failed to open zip file: %s", err)
			continue
		}
		defer f.Close()

		d, err := ioutil.ReadAll(f)
		if err != nil {
			t.Errorf("Failed to read file contents: %s", err)
			continue
		}

		b := bytes.NewReader(d)
		c, err := NewClass(b)
		if err != nil {
			t.Errorf("Failed to create Class struct: %s", err)
			continue
		}

		if len(v) == 0 {
			t.Log(c.String())
			continue
		}

		if d, err := diff(v, []byte(c.String())); err != nil {
			t.Error(err)
		} else if len(d) != 0 {
			t.Error(string(d))
		}
	}
}
