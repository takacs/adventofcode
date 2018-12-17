import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;
import java.util.Scanner;

public class Legnagyobbhelyi {

	public static boolean isNum(String str) {
		try {
			int d = Integer.parseInt(str);
			return true;
		} catch (NumberFormatException nfe) {
			return false;
		}
	}
	
	public static void main(String[] args) {
		Scanner sc = new Scanner(System.in);
		Map<Integer, Integer> dig = new HashMap<Integer, Integer>();
		for(int i = 0; i < 10; i++) {
			dig.put(i, 0);
		}
		
		ArrayList<String> nums = new ArrayList<String>();
		String line;
		while(sc.hasNextLine()) {
			line = sc.nextLine();
			String[] strings = line.split("\\s+");
			for (String string : strings) {
				if (isNum(string)) { nums.add(string); }
			}
		}
		
		for (int i = 0; i < nums.size(); i++) {
			char c = nums.get(i).charAt(0);
			if(c == '-') { c = nums.get(i).charAt(1); }
			String s = Character.toString(c);
			int n = Integer.parseInt(s);
			dig.put(n, dig.get(n) + 1);
		}
		
		for(int i1 = 0; i1 < 10; i1++) {
			System.out.println(i1 + ": " + dig.get(i1) + " db");
		}
		
		
		sc.close();
	}

}
