import java.io.*;
import java.util.Arrays;

public class Legkisebb {
    public static void main(String[] args) {
        try (BufferedReader br = new BufferedReader(new InputStreamReader(System.in))) {
            br.lines()
                    .map(line -> line.split("\\s"))
                    .forEach(tokens -> Arrays.stream(tokens)
                            .filter(s -> !s.isEmpty())
                            .mapToInt(Integer::parseInt)
                            .min()
                            .ifPresentOrElse(System.out::println, () -> System.out.println("unknown"))
                    );
        } catch (IOException e) {
        }
    }
}
