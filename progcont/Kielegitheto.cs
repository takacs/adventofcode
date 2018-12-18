using System;

namespace Kielegitheto
{
    class Program
    {
        static void Main(string[] args)
        {
            string line;
            string op = "&";
            while ((line = Console.ReadLine()) != null)
            {
                string[] l = line.Split(' ');
                if (l.Length == 3) { op = l[1]; Console.WriteLine(line);  }
                else {
                    int[] n = Array.ConvertAll<string, int>(l, int.Parse);
                    switch (op)
                    {
                        case "&":
                            if (n[0] == 1 && n[1] == 1) { Console.WriteLine(1); }
                            else { Console.WriteLine(0); }
                            break;
                        case "|":
                            if (n[0] == 1 || n[1] == 1) { Console.WriteLine(1); }
                            else { Console.WriteLine(0); }
                            break;
                        case "=":
                            if (n[0] == 1 && n[1] == 1) { Console.WriteLine(1); }
                            else if (n[0] == 0 && n[1] == 0) { Console.WriteLine(1); }
                            else { Console.WriteLine(0); }
                            break;
                        case ">":
                            if (n[0] == 1 && n[1] == 0) { Console.WriteLine(0); }
                            else { Console.WriteLine(1); }
                            break;
                    }
                 
                }
            }
        }
    }
}
