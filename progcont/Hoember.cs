using System;
using System.Collections.Generic;

namespace Hoember
{
    class Program
    {
        static bool IsProp(int a, int b, int c)
        {
            double top = (double)a / (double)b;
            double bot = (double)b / (double)c;
            if (top == bot) { return true; }
            return false;
        }

        static bool IsStab(int a, int b, int c)
        {
            if(a>b && b > c) { return true; }
            return false;
        }
        static void Main(string[] args)
        {
            string line;
            int a, b, c;
            while ((line = Console.ReadLine()) != null){
                string[] nums = line.Split(' ');
                Console.Write(line + " ");
                a = Int32.Parse(nums[0]);
                b = Int32.Parse(nums[1]);
                c = Int32.Parse(nums[2]);
                
                if (IsStab(a,b,c) && IsProp(a, b, c)) { Console.WriteLine("STABLE AND PROPORTIONAL"); }
                if (IsStab(a, b, c) && !IsProp(a, b, c)) { Console.WriteLine("ONLY STABLE"); }
                if (!IsStab(a, b, c) && IsProp(a, b, c)) { Console.WriteLine("ONLY PROPORTIONAL"); }
                if (!IsStab(a, b, c) && !IsProp(a, b, c)) { Console.WriteLine("NEITHER STABLE NOR PROPORTIONAL"); }

            } 
        }
    }
}
