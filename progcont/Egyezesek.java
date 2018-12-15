import java.util.Scanner;

public class Egyezesek {

	public static void main(String[] args) {
		Scanner scanner = new Scanner(System.in);
		String line, a, b;
		while(!(line = scanner.nextLine()).equals("END")) {
			int count = 0;
			a = line.split(" ")[0];
			b = line.split(" ")[1];
			for(int i = 0; i < b.length(); i++) {
				try{
					if (b.charAt(i) == a.charAt(i)){
						count++;
						}
					} catch (StringIndexOutOfBoundsException e) {
					}
				}
			System.out.println(a+"-"+b+": " + count);
			}
		scanner.close();	
		}
	}
