import java.util.Arrays;
import java.util.Scanner;

class Bolt implements Comparable<Bolt>{
	String name;
	String cat;
	int val;
	
	public Bolt(String name, String cat, int val) {
		super();
		this.name = name;
		this.cat = cat;
		this.val = val;
	}

	@Override
	public int compareTo(Bolt o) {
		if ( this.val > o.val) {
			return -1;
		}
		else if ( this.val < o.val) { return 1; }
		else {
			int ncmp = this.name.compareTo(o.name);
			if(ncmp > 0) {
				return 1;
			} else if ( ncmp < 0) {
				return -1;
			} else {
				int tcmp = this.cat.compareTo(o.cat);
				if(tcmp > 0) {
					return 1;
				} else if ( tcmp < 0) {
					return -1;
				}
			}
		}
		return 0;
	}

	@Override
	public String toString() {
		return cat + " (" + name + "): " + val + "%";
	}
	
}

class Bolt2 implements Comparable<Bolt2>{
	String name;
	String cat;
	int val;
	
	public Bolt2(String name, String cat, int val) {
		super();
		this.name = name;
		this.cat = cat;
		this.val = val;
	}

	@Override
	public int compareTo(Bolt2 o) {
		int ncmp = this.name.compareTo(o.name);
		if ( ncmp > 0) {
			return 1;
		}
		else if ( ncmp < 0) { return -11; }
		else {
			if(this.val > o.val) {
				return -1;
			}
			else if (this.val < o.val) {return 1;}
			else {
				int tcmp = this.cat.compareTo(o.cat);
				if(tcmp > 0) {
					return 1;
				} else if ( tcmp < 0) {
					return -1;
				}
			}
		}
		return 0;
	}

	@Override
	public String toString() {
		return cat + " (" + name + "): " + val + "%";
	}
	
}

public class Kuponok {

	public static void main(String[] args) {
		Scanner sc = new Scanner(System.in);
		int n = Integer.parseInt(sc.nextLine());
		Bolt[] boltok = new Bolt[n];
		Bolt2[] boltok2 = new Bolt2[n];
		for (int i = 0; i < n; i++) {
			String[] line = sc.nextLine().split(";");
			boltok[i] = new Bolt(line[0], line[1], Integer.parseInt(line[2]));
			boltok2[i] = new Bolt2(line[0], line[1], Integer.parseInt(line[2]));
		}
		Arrays.sort(boltok);
		Arrays.sort(boltok2);
		
		for (int i = 0; i < boltok.length; i++) {
			System.out.println(boltok[i]);
		}
		System.out.println();
		for (int i = 0; i < boltok2.length; i++) {
			System.out.println(boltok2[i]);
		}
	}

}
