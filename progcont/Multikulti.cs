using System;
using System.Collections.Generic;

namespace Multikulti
{
    class Program
    {
        static void Main(string[] args)
        {
            Dictionary<string, string> szotar = new Dictionary<string, string>();
            string line;
            while ((line = Console.ReadLine()) != null)
            {
                string[] l = line.Split(' ');
                if (l.Length > 1)
                {
                    if (szotar.ContainsKey(l[0]))
                    {
                        if (szotar[l[0]].Length > l[0].Length) { szotar[l[0]] = l[1]; }
                    }
                    else { szotar[l[0]] = l[1]; }
                }
                else
                {
                    if (szotar.ContainsKey(l[0])) { Console.WriteLine(szotar[l[0]]); }
                    else { Console.WriteLine("unknown"); }
                }
            }
        }
    }
}