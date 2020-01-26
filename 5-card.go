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
   return 1;                        /* return 1 if all card belong to same suit */
}

/* Function for hand value Straight*/
func Straight() int
{
	if play[0][0]=='2' && play[1][0]=='3' && play[2][0]=='4' && play[3][0]=='5' && play[4][0]=='A' ||
	   play[0][0]=='2' && play[1][0]=='3' && play[2][0]=='4' && play[3][0]=='5' && play[4][0]=='6' ||
	   play[0][0]=='3' && play[1][0]=='4' && play[2][0]=='5' && play[3][0]=='6' && play[4][0]=='7' ||
	   play[0][0]=='4' && play[1][0]=='5' && play[2][0]=='6' && play[3][0]=='7' && play[4][0]=='8' ||
	   play[0][0]=='5' && play[1][0]=='6' && play[2][0]=='7' && play[3][0]=='8' && play[4][0]=='9' ||
	   play[0][0]=='6' && play[1][0]=='7' && play[2][0]=='8' && play[3][0]=='9' && play[4][0]=='T' ||
	   play[0][0]=='7' && play[1][0]=='8' && play[2][0]=='9' && play[3][0]=='J' && play[4][0]=='T' ||
	   play[0][0]=='8' && play[1][0]=='9' && play[2][0]=='J' && play[3][0]=='Q' && play[4][0]=='T' ||
	   play[0][0]=='9' && play[1][0]=='J' && play[2][0]=='K' && play[3][0]=='Q' && play[4][0]=='T' ||
	   play[0][0]=='A' && play[1][0]=='J' && play[2][0]=='K' && play[3][0]=='Q' && play[4][0]=='T'
	(
		return 1;       /* if player have card of same sequence but of different suit then user has straight hand*/

	)

	return 0;
}


func main() {

}