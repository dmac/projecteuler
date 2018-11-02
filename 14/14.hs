import Data.Map (Map)
import qualified Data.Map as Map

main :: IO ()
main = putStrLn $ show $ longestChain [1..1000000]

longestChain :: [Int] -> (Int, Int)
longestChain nums = longestChain' nums Map.empty

longestChain' :: [Int] -> (Map Int Int) -> (Int, Int)
longestChain' [] _ = (0, 0)
longestChain' (x:xs) lengths =
  let (xLength, newLengths) = fullChainLength x lengths
      thisBest = (x, xLength)
      thatBest = longestChain' xs newLengths
  in  if (snd thisBest) > (snd thatBest)
      then thisBest else thatBest

fullChainLength :: Int -> (Map Int Int) -> (Int, (Map Int Int))
fullChainLength 1 lengths = (1, lengths)
fullChainLength n lengths =
  case Map.lookup n lengths of
    Nothing -> let (theLength, newLengths) = fullChainLength (nextInChain n) lengths
               in  (theLength + 1, Map.insert n (theLength + 1) newLengths)
    Just l  -> (l, lengths)

fullChain :: Int -> [Int]
fullChain 1 = [1]
fullChain n = n : fullChain (nextInChain n)

nextInChain :: Int -> Int
nextInChain n
  | even n = n `div` 2
  | odd n  = 3 * n + 1
