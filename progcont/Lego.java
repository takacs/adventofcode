import java.util.Arrays;
import java.util.Scanner;

class Keszlet implements Comparable<Keszlet>{
	int kod;
	String szett;
	String tema;
	int db;
	
	public Keszlet(int kod, String szett, String tema, int db) {
		super();
		this.kod = kod;
		this.szett = szett;
		this.tema = tema;
		this.db = db;
	}

	@Override
	public int compareTo(Keszlet o) {
		if(this.db > o.db) { return -1; }
		else if (this.db < o.db) { return 1; }
		else {
			int tcmp = this.tema.compareTo(o.tema);
			if(tcmp > 0) { return 1; }
			else if (tcmp < 0) { return -1; }
			else {
				int ncmp = this.szett.compareTo(o.szett);
				if(ncmp > 0) { return 1; }
				else if (ncmp < 0) { return -1; }
				else {
					if(this.kod > o.kod) { return 1; }
					else if(this.kod < o.kod) { return -1; }
				}
			}
		}
		return 0;
	}
	@Override
	public String toString() {
		return this.szett + " ("+this.kod + "): " + this.db + " - " + this.tema;
	}
}

class Keszlet2 implements Comparable<Keszlet2>{
	int kod;
	String szett;
	String tema;
	int db;
	
	public Keszlet2(int kod, String szett, String tema, int db) {
		super();
		this.kod = kod;
		this.szett = szett;
		this.tema = tema;
		this.db = db;
	}

	@Override
	public int compareTo(Keszlet2 o) {
		int tcmp = this.tema.compareTo(o.tema);
		if (tcmp > 0) {return 1;}
		else if (tcmp < 0) {return -1;}
		else {
			int ncmp = this.szett.compareTo(o.szett);
			if (ncmp > 0) { return 1; }
			else if (ncmp < 0) { return -1; }
			else {
				if (this.kod > o.kod) { return 1; }
				else if (this.kod < o.kod ) { return -1; }
			}
		}
		return 0;
	}
	@Override
	public String toString() {
		return this.szett + " ("+this.kod + "): " + this.db + " - " + this.tema;
	}
}

public class Lego {

	public static void main(String[] args) {
		Scanner sc = new Scanner(System.in);
		int n = Integer.parseInt(sc.nextLine());
		Keszlet[] ksz = new Keszlet[n];
		Keszlet2[] ksz2 = new Keszlet2[n];
		for (int i = 0; i < n; i++) {
			String l = sc.nextLine();
			String[] line = l.split(";");
			ksz[i] = new Keszlet(Integer.parseInt(line[0]), line[1], line[2], Integer.parseInt(line[3]));
			ksz2[i] = new Keszlet2(Integer.parseInt(line[0]), line[1], line[2], Integer.parseInt(line[3]));

		}
		Arrays.sort(ksz);
		Arrays.sort(ksz2);
		
		for(int i = 0 ; i<ksz.length; i++) {
			System.out.println(ksz[i]);
		}
		System.out.println();
		for(int i = 0 ; i<ksz2.length; i++) {
			System.out.println(ksz2[i]);
		}
		sc.close();
	}

}
