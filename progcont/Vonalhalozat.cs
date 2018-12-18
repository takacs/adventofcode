using System;
using System.Collections.Generic;

namespace Vonalhalozat
{
    class Program
    {
        static void Main(string[] args)
        {
            List<string> vegallomas = new List<string>();
            List<string> koztes = new List<string>();
            for(int i = 0; i < 3; i++)
            {
                System.IO.StreamReader file;
                try { file = new System.IO.StreamReader(args[i]); }
                catch (Exception) { continue; }

                string line;
                line = file.ReadLine();
                if (!vegallomas.Contains(line)) { vegallomas.Add(line); }
                while((line = file.ReadLine()) != null){
                    if (line.StartsWith("/")) { continue; }
                    if(file.Peek() >= 0)
                    {
                        if (!koztes.Contains(line)) { koztes.Add(line); }
                        
                    }
                    else
                    {
                        if (!vegallomas.Contains(line)) { vegallomas.Add(line); }
                    }
                }
            }

            koztes.Sort();
            foreach (string s in koztes)
            {
                if (vegallomas.Contains(s)) { Console.WriteLine(s);  }
            }
        }
    }
}
