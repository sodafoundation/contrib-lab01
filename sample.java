// PALINDROME NUMBER

import java.util.*;
class sample
{
    public static void main(String args[])
    {
        Scanner mg = new Scanner (System.in);
        int num,num1,d,rev=0;
        System.out.println("ENTER THE NUMBER");
        num=mg.nextInt();
        num1=num;
        while(num1>0)
        {
            d=num1%10;
            rev=rev*10+d;
            num1/=10;
        }
        if(num==rev)
        System.out.println("PALINDROME NUMBER");
        else
        System.out.println("NOT PALINDROME NUMBER");

    }
}

