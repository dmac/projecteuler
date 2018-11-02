import Data.Map (Map)
import qualified Data.Map as Map

digitWords = Map.fromList [(1, "one")
                          ,(2, "two")
                          ,(3, "three")
                          ,(4, "four")
                          ,(5, "five")
                          ,(6, "six")
                          ,(7, "seven")
                          ,(8, "eight")
                          ,(9, "nine")
                          ]

tensWords = Map.fromList [(2, "twenty")
                         ,(3, "thirty")
                         ,(4, "forty")
                         ,(5, "fifty")
                         ,(6, "sixty")
                         ,(7, "seventy")
                         ,(8, "eighty")
                         ,(9, "ninety")
                         ]

numberToWord :: Int -> String
numberToWord num = undefined

digitToWord :: Int -> Int -> String
digitToWord digit place =
  case place of
    4 -> (Map.findWithDefault "" digit digitWords) ++ "thousand"
    3 -> (Map.findWithDefault "" digit digitWords) ++ "hundred"
    2 -> (Map.findWithDefault "" digit tensWords)
    1 -> (Map.findWithDefault "" digit digitWords)

