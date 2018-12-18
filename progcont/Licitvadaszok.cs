using System;
using System.Collections.Generic;

namespace Licitvadaszok
{
    class Program
    {
        static void Main(string[] args)
        {
            Dictionary<string, int> ppl = new Dictionary<string, int>();
            string line;
            while((line = Console.ReadLine()) != null)
            {
                string[] l = line.Split(';');
                if (l.Length == 2)
                {
                    if (ppl.ContainsKey(l[0]))
                    {
                        int curr = ppl[l[0]];
                        if (curr < Int32.Parse(l[1])) { ppl[l[0]] = Int32.Parse(l[1]); }
                    }
                    else
                    { 
                        ppl.Add(l[0], Int32.Parse(l[1]));
                    }
                }
                else
                {
                    if (ppl.ContainsKey(l[0]))
                    {
                        Console.WriteLine(ppl[l[0]]);
                    }
                    else { Console.WriteLine("unknown");  }
                   
                }
            }
        }
    }
}
