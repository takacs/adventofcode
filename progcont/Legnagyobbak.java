import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.util.Arrays;

public class Legnagyobbak {

	public static void main(String[] args) {
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
		br.lines()
			.map(line -> line.split("\\s"))
			.forEach(tokens -> Arrays.stream(tokens)
					.filter(s -> !s.isEmpty())
					.mapToInt(Integer::parseInt)
					.filter(i -> i%2 == 0)
					.max()
					.ifPresentOrElse(System.out::println,() -> System.out.println("unknown")));
	}

}
