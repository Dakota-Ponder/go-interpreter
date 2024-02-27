// the lexer takes the source code as input and returns a stream of tokens that 
// represent the source code. It will fo through the input and output the next 
// token it recognizes.

package lexer

// struct for the lexer 
type Lexer struct {
	input 		 string // the input source code
	position	 int    // the current position in the input
	readPosition int    // the next position in the input
	ch 			 byte   // the current character under examination
}

// function that creates a new lexer
// returns a pointer to the lexer object 
func New(input string) *Lexer {
	l := &Lexer{input: input} // use a struct literal to set the input field of the lexer to the input argument
	return l // return the pointer to the lexer 
}

// helper function 
func (l *Lexer) readChar(){
	// check if the readPostition is greater than or equal to the length 
	// of the input string
	// in other words, check if we have reached the end of the input 
	if l.readPosition >= len(l.input){
		// if it is, set the character to 0 
		l.ch = 0 // ASCII code for "NUL"
	} else {
		// if it is not the end of the input
		// , set the character to the character at the readPosition
		l.ch = l.input[l.readPosition]
	}
	// set the position to the last place we read
	l.position = l.readPosition
	
	// increment the readPosition to where we will be next 
	l.readPosition += 1
}