-- Jan 1 1901 was a Tuesday

-- 0 Sunday
-- 1 Monday
-- 2 Tuesday
-- 3 Wednesday
-- 4 Thursday
-- 5 Friday
-- 6 Saturday

main :: IO ()
main =
    let days = [(year, month, day) | year <- [1901..2000],
                                     month <- [1..12],
                                     day <- [1..daysInMonth month year]]
        daysWithIndex = zip days ([2..] :: [Int])
    in print $ length $ filter id $ map (\((_, _, day), i) -> i `mod` 7 == 0 && day == 1) daysWithIndex

daysInMonth :: Int -> Int -> Int
daysInMonth month year
  | month `elem` [1, 3, 5, 7, 8, 10, 12] = 31
  | month `elem` [4, 6, 9, 11] = 30
  | month == 2 = if year `mod` 4 == 0 || (year `mod` 100 == 0 && year `mod` 400 == 0)
                   then 29
                   else 28
  | otherwise = error $ "Invalid month passed to daysInMonth (must be in [1..12]): " ++ show month
