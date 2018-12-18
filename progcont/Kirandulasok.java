import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.HashMap;
import java.util.Map;

public class Kirandulasok {

	public static void main(String[] args) throws IOException {
		BufferedReader br = new BufferedReader(new FileReader(args[0]));
		String line;
		Map<String, Integer> tourist = new HashMap<String, Integer>();
		while((line = br.readLine()) != null) {
			String[] l = line.split(":");
			l = l[1].split(";");
			for(int i = 0; i<l.length;i++) {
				try {
					tourist.put(l[i], tourist.get(l[i]) + 1);
				}catch(NullPointerException n) {
					tourist.put(l[i], 1);
				}
			}
		}
		
		int max = tourist.entrySet().stream()
                .mapToInt(Map.Entry::getValue)
                .max()
                .getAsInt();
		
		tourist.entrySet().stream().filter(e-> e.getValue() == max).map(Map.Entry::getKey).sorted().forEach(System.out::println);
		br.close();
	}	

}
