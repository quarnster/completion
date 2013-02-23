public class Basic {

	interface Callable {
		void call();
	}

	private Callable mCallable = new Callable() {
		void call() { }
	};

	private String mText;

	public String getText() {
		return mText;
	}

	private int mPos;

}