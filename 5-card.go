package main

import "fmt"


/* Group all (related values )hand value like enum in c++*/
const 
(
    HIGHEST_CARD = 0
	ONE_PAIR
	TWO_PAIRS
	THREE_OF_A_KIND
	STRAIGHT
	FLUSH
	FULL_HOUSE
	FOUR_OF_A_KIND
	STRAIGHT_FLUSH
)

/* Declration and initialization of a pointer array */
var *hands[] char

hands[0] = "highest-card"
hands[1] = "one-pair"
hands[2] = "two-pairs"
hands[3] = "three-of-a-kind"
hands[4] = "straight"
hands[5] = "flush"
hands[6] = "full-house"
hands[7] = "four-of-a-kind"
hands[8] = "straight-flush"

/* declaration of arrays of hand card , deck and card which is currently in play */
var hand[5][3]char
var hand[5][3]char
var hand[5][3]char

/* Function for hand value Flush */
func Flush() int 
{
   var i int;
   var c char;
   c = play[0][1];                  /* c stores suit to which card belongs */

   for i:= 1 ; i<5 ; i++
   {
	   if c!=play[i][1]              /* compares the next card's suit */
	   {
		   return 0;
	   }
   }
   return 1;                        /* return 1 if all card belong to same suit
}





func main() {

}