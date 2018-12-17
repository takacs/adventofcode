import java.util.Scanner;

public class Kulonbsegek {

	public static void main(String[] args) {
		Scanner scanner = new Scanner(System.in);
		String line, a, b;
		while(!(line = scanner.nextLine()).equals("END")) {
			int count = 0;
			a = line.split(" ")[0];
			b = line.split(" ")[1];
			int max;
			if(a.length() > b.length()) { max = a.length(); }
			else { max = b.length();}
			for(int i = 0; i < max; i++) {
				try{
					if (b.charAt(i) != a.charAt(i)){
						count++;
						}
					} catch (StringIndexOutOfBoundsException e) {
						count++;
					}
				}
			System.out.println(a+"-"+b+": " + count);
			}
		scanner.close();	
		}
	}
