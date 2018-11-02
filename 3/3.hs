import qualified Data.Set as Set

main :: IO ()
main = print $ maximum $ primeFactors 600851475143

primeFactors :: Int -> [Int]
primeFactors n =
    Set.toList $
    foldr (\f set -> if n `mod` f == 0
                       then let subFactors = primeFactors f
                            in if Set.null $ Set.fromList subFactors
                                 then Set.insert f set
                                 else set `Set.union` Set.fromList subFactors
                       else set)
          Set.empty
          [2 .. (floor . (sqrt::Double->Double) . fromIntegral) n]
