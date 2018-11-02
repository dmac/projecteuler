main :: IO ()
main = print $ sum $ filter (\n -> n `mod` 2 == 0) $ takeWhile (< 4000000) fibs

fibs :: [Int]
fibs = 1 : 2 : zipWith (+) fibs (tail fibs)
