import Data.Map (Map)
import qualified Data.Map as Map
import Data.Maybe (fromMaybe)
import List (sortBy)
import Data.Ord (comparing)


main :: IO ()
main = printLengths [200000..300000] []

printLengths :: [Int] -> [Int] -> IO ()
printLengths [] _ = return ()
printLengths (n:ns) lengths = do
  if n `mod` 100000 == 0 && (length lengths > 0)
    then putStrLn $ (show n) ++ ": " ++ (show . maximum) lengths
    else return ()
  printLengths ns $ (fullChainLengthSimple n) : lengths
  return ()

{-main :: IO ()-}
{-main =-}
  {-putStrLn $ show $ last $ sortBy (comparing snd) $ Map.toList $ lengthMap-}
  {-where lengthMap = foldr fullChainLength Map.empty [1..1000000]-}

fullChainLength :: Int -> (Map Int Int) -> (Map Int Int)
fullChainLength 1 lengths = Map.insert 1 1 lengths
fullChainLength n lengths =
  case Map.lookup n lengths of
    Just l  -> lengths
    Nothing -> let next = nextInChain n
                   newLengths = fullChainLength next lengths
                   nextLength = fromMaybe (error "blah") (Map.lookup next newLengths)
               in  Map.insert n (1 + nextLength) newLengths

fullChainLengthSimple :: Int -> Int
fullChainLengthSimple 1 = 1
fullChainLengthSimple n = 1 + fullChainLengthSimple (nextInChain n)

nextInChain :: Int -> Int
nextInChain n
  | even n = n `div` 2
  | odd n  = 3 * n + 1
