using System;

namespace Hullamvasut
{
    class Vasut : IComparable<Vasut>
    {
        string n, c;
        int h, t;

        public Vasut(string n, string c, int h, int t)
        {
            this.n = n;
            this.c = c;
            this.h = h;
            this.t = t;
        }

        public int CompareTo(Vasut o)
        {
            if(this.t>o.t) { return 1; }
            else if(this.t<o.t) { return -1; }
            else
            {
                if (this.h > o.h) { return -1; }
                else if(this.h < o.h) { return 1; }
                else
                {
                    int ncmp = this.n.CompareTo(o.n);
                    if (ncmp > 0) { return 1; }
                    else { return -1; }
                }
            }
        }

        public override string ToString()
        {
            return this.n + " (" + this.c + "): " + Convert.ToString(this.t);
        }
    }
    class Program
    {
        static void Main(string[] args)
        {
            int n = Int32.Parse(Console.ReadLine());
            Vasut[] v = new Vasut[n];
            for(int i = 0; i < n; i++)
            {
                string[] line = Console.ReadLine().Split(';');
                v[i] = new Vasut(line[0], line[1], Convert.ToInt32(line[2]), Convert.ToInt32(line[3]));
            }

            Array.Sort(v);

            for(int i = 0; i < n; i++){
                Console.WriteLine(v[i]);
            }
            Console.ReadLine();
        }
    }
}
