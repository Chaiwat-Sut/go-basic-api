var banknote = [1000,500,100,50,20,10,5,2,1]

function sell(price,pay) {
    let result = []
    let c = pay - price
    for (let note of banknote) {
        if (c >= note ) {
            noteAmount = Math.floor(c / note)
            c = c - (noteAmount * note)
            result.push({banknote: note,amount: noteAmount})
        }
    }
    return (result)
}

console.log(sell(20,1000))