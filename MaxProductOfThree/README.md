# My Thought Process

## Goal
The goal here is to find three maximal numbers and two minimal.

And in order to find the maximal product, one  only need to find out between Max * Max * Max and Max * Min * Min, which is bigger.

## Why

Consider all possible combinations of three numbers. 

There four possibilities, which are 

1. + + + 
2. - - - 
3. - - +
4. - + + 


And obviously case 4 makes no sense. Given any input array bigger than 3, you will pick a combination from 1,2,3, but not 4. 

Case 2 only make sense when all integers are negative in the input array, and under that case, case 1 is actually the same as case 2 (after you sorted the array)

So the only left over combinations are case 1 and case 3, which we really can't determine the result but looking at them. So we have to calculate them. 

One don't need to calculate when input_array.length <= 4, since the result is obvious by only looking at the signs. 

In others words, if there is only 1 negative integer, you don't care, and  + * + * + is naturally bigger than - * + * +.

If there are 2 negative integers (or more), you calculate it, since multiplying two negatives may make a big positive number.

If all integers are negative, you will end up in case 2 after the comparison, if there is one or more than one positive integer, you will end up either case 1 or case 3. Only two positive integers will not make sense as case 4 shows.

In short, only case 1 or case 3 is a possible solution.