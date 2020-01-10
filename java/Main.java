public class Main {
	public static void main(String[] args) {
		while (true) {
			foo(10);
		}
	}

	private static void foo(int i) {
		if (i <= 0) {
			return;
		}
		foo(i - 1);
	}
}
