package testdata;

import java.lang.annotation.ElementType;
import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;
import java.lang.annotation.Target;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.concurrent.Future;

public class Basic {

	static int sVal;

	static {
		sVal = 2 * 3;
	}

	public static int getVal() {
		return sVal;
	}

	@Target(ElementType.METHOD)
	@Retention(RetentionPolicy.CLASS)
	public @interface Special { }

	interface Callable {
		void call();
	}

	private Callable mCallable = new Callable() {
		public void call() { }
	};

	private String mText;

	private Future<?> mFuture;

	private Map<Integer, List<String>> mMap = new HashMap<Integer, List<String>>();

	public String getText() {
		char c = '}';
		mText = "{hello};";
		return mText;
	}

	@Special
	public void setText(String text) {
		mText = text;
	}

	private int mPos;

}
