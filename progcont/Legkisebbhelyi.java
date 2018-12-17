import java.io.*;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;
import java.util.Scanner;


public class Legkisebbhelyi {
	
	public static boolean isNumeric(String str) {
		
		try {
			int d = Integer.parseInt(str);
			return true;
		} catch(NumberFormatException nfe) {
			return false;
		}
	}
	public static void main(String[] args) throws IOException {
		Scanner sc = new Scanner(System.in);
		String line;
		Map<Integer, Integer> map = new HashMap<Integer, Integer>();
		for(int i = 0; i < 10; i++) {
			map.put(i , 0);
		}
		ArrayList<Integer> nums = new ArrayList<Integer>();
		while(sc.hasNextLine()) {
				line = sc.nextLine();
				String[] strings = line.trim().split("\\s+");

				for (int i = 0; i < strings.length; i++) {
					if (isNumeric(strings[i])) {
						nums.add(Integer.parseInt(strings[i]));
					}
				}
				
			} // while
		
		for (int i = 0; i < nums.size(); i++) {
			int low = nums.get(i)%10;
			if (low <0) {low *= -1;}
			map.put(low, map.get(low) + 1);
		}
		
		for(int i = 0; i < 10; i++) {
			System.out.println(i + ": " + map.get(i) + " db");
		}
		
		}
    
				
}
