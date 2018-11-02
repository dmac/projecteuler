import Data.Maybe (catMaybes)
import Data.List (find)

main :: IO ()
main = putStrLn $ show $ find (\n -> (length . factors) n > 500) (map triangleN [1..])

triangleN :: Integer -> Integer
triangleN 1 = 1
triangleN n = n + triangleN (n - 1)

factors n =
  let candidates = [1..floor (sqrt (fromInteger n))]
  in concat $ catMaybes $ map (\x -> if n `mod` x == 0
                                     then Just [x, n `div` x]
                                     else Nothing) candidates

