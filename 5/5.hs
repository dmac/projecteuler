import Data.List (find)
import Data.Maybe (fromMaybe)

main :: IO ()
main = do
  putStrLn $ show $ smallestDivisibleByEach [1..20]

smallestDivisibleByEach :: [Integer] -> Integer
smallestDivisibleByEach nums = fromMaybe 0 $ find (allDivide nums) [1..(productOfAll [1..20])]

allDivide :: [Integer] -> Integer -> Bool
allDivide nums target = and $ map (\n -> n `divides` target) nums

divides :: Integer -> Integer -> Bool
a `divides` b = b `mod` a == 0

productOfAll :: [Integer] -> Integer
productOfAll nums = foldr (\num acc -> num * acc) 1 nums
