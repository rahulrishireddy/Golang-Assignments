package HouseCup

import (
	"fmt"
	"math/rand"
	"time"
)

func Randomnumbers() {
	var scores [160]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 160; i++ {
		scores[i] = rand.Intn(100)

	}
	// fmt.Println(scores) // [0 28 27 62 63]// Gives you all the 160 random numbers generated.

	var Gryffindor []int = scores[:40]     // slicing for gryffindor
	var Slytherin []int = scores[40:80]    // slicing for slytherin
	var Ravenclaw []int = scores[80:120]   //slicing for ravenclaw
	var Hufflepuff []int = scores[120:160] // slicing for Hufflepuff

	Gryffindorresult := 0
	for _, v := range Gryffindor {
		Gryffindorresult += v
	}
	Slytherinresult := 0
	for _, v := range Slytherin {
		Slytherinresult += v
	}

	Ravenclawresult := 0
	for _, v := range Ravenclaw {
		Ravenclawresult += v
	}

	Hufflepuffresult := 0
	for _, v := range Hufflepuff {
		Hufflepuffresult += v
	}

	fmt.Println(Gryffindor) // Gives you the random numbers generated for gryffindor
	fmt.Println(Slytherin)  // Gives you the random numbers generated for slytherin
	fmt.Println(Ravenclaw)  //Gives you the random numbers generated for Ravenclaw
	fmt.Println(Hufflepuff) //Gives you the random numbers generated for Hufflepuff

	fmt.Println("Gryffindorresult is", Gryffindorresult)
	fmt.Println("Slytherinresult is", Slytherinresult)
	fmt.Println("Ravenclawresult is", Ravenclawresult)
	fmt.Println("Hufflepuffresult is", Hufflepuffresult)

	fmt.Println(Slytherin[9], Slytherin[24], Slytherin[28])

	//printed the scores of slytherin with id number (10,25,29)
	Slytherin[9] = Slytherin[9] - 2
	Slytherin[24] = Slytherin[24] - 2
	Slytherin[28] = Slytherin[28] - 2
	// In the above code we decremented 2 points to students in slytherin with id (10,25,29)
	fmt.Println(Slytherin[9], Slytherin[24], Slytherin[28])

	fmt.Println(Gryffindor[8], Gryffindor[26], Gryffindor[14])
	//printed the scores of Harry, Hermoine, Ron  with id number (9,27,15)
	Gryffindor[8] = Gryffindor[8] + 4
	Gryffindor[26] = Gryffindor[26] + 4
	Gryffindor[14] = Gryffindor[14] + 4
	// In the above code we increased 4 points to students in slytherin with id (9,27,15)
	fmt.Println(Gryffindor[8], Gryffindor[26], Gryffindor[14])

	for i := 0; i <= 39; i++ { // Ravenclaw lost the “quiddich” game(game played with brooms)  and each student in the group lost 1 point

		Ravenclaw[i] = Ravenclaw[i] - 1

	}
	fmt.Println("Ravenclaw", Ravenclaw)

	for i := 0; i <= 39; i++ { //Hufflepuff on the other hand stood 3rd place in “quiddich” game and each student in the group gain 2 points

		Hufflepuff[i] = Hufflepuff[i] + 1

	}
	fmt.Println("Hufflepuff", Hufflepuff)

	GDupdatedresult := 0
	for _, v := range Gryffindor {
		GDupdatedresult += v
	}
	SYupdatedresult := 0
	for _, v := range Slytherin {
		SYupdatedresult += v
	}

	RCupdatedresult := 0
	for _, v := range Ravenclaw {
		RCupdatedresult += v
	}

	HPupdatedresult := 0
	for _, v := range Hufflepuff {
		HPupdatedresult += v
	}

	fmt.Println(Gryffindor) // Gives you the random numbers generated for gryffindor
	fmt.Println(Slytherin)  // Gives you the random numbers generated for slytherin
	fmt.Println(Ravenclaw)  //Gives you the random numbers generated for Ravenclaw
	fmt.Println(Hufflepuff) //Gives you the random numbers generated for Hufflepuff

	fmt.Println("Gryffindorresult is", GDupdatedresult)
	fmt.Println("Slytherinresult is", SYupdatedresult)
	fmt.Println("Ravenclawresult is", RCupdatedresult)
	fmt.Println("Hufflepuffresult is", HPupdatedresult)

	GDaverage := float64(GDupdatedresult) / 40 // Print Average points scored by group Gryffindor
	fmt.Println("Gryffindoravg is", GDaverage)

	SYaverage := float64(SYupdatedresult) / 40 // Print Average points scored by group Slytherin
	fmt.Println("Slytherinavg is", SYaverage)

	RCaverage := float64(RCupdatedresult) / 40 // Print Average points scored by group Ravenclaw
	fmt.Println("Ravenclawavg is", RCaverage)

	HPaverage := float64(HPupdatedresult) / 40 // Print Average points scored by group Hufflepuff
	fmt.Println("Hufflepuffavg is", HPaverage)

	fmt.Println("GRYFFINDOR WINS THE HOUSE CUP ")

}
