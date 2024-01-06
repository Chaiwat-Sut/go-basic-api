fizzBuzz :: Int -> String
fizzBuzz 3 = "Fizz"
fizzBuzz 5 = "Buzz"
fizzBuzz n
  | n `mod` 3 == 0 = fizzBuzz 3
  | n `mod` 5 == 0 = fizzBuzz 5
  | otherwise = show n

main :: IO ()
main = do
  x <- getLine
  print (fizzBuzz (read x))