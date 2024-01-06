// println("Hello, World!")
val FizzBuzz = { n : Int ->
     when {
        n % 15 == 0 -> "FizzBuzz"
        n % 5 == 0 -> "Buzz"
        n % 3 == 0 -> "Fizz"
        else -> n.toString()
    }
}
println(FizzBuzz(15))
