
open System

let parseInt str = str |> int

// parses into Some or catches exception and resolves to None
let maybeParseInt (str: string) =
    try 
        str |> int |> Some
    with _ -> None

// Computation Expression
type ParseComputationWithLog() =
    member this.Bind(m,f) =
        printf "Value is %A" m
        Option.bind f m
    member this.Return(x) = Some(x)

let pcwl = new ParseComputationWithLog()

let workflow s1 s2 s3 =
    pcwl {
        let! i1 = s1 |> maybeParseInt
        let! i2 = s2 |> maybeParseInt
        let! i3 = s3 |> maybeParseInt
        return i1 + i2 + i3
    }

let good = workflow "12" "3" "2"
printf "%A" good

let bad = workflow "12" "xyz" "2"
printf "%A" bad

// Output so far:
// Value is 12 (from good)
// Value is 3 (from good)
// Value is 2 (from good)
// 17 -> Result of good
// Value is 12 (from bad)
// -> Nothing else because xyz is invalid and results in None

let (>>=) m f = f m

let strAdd str i = parseInt str >>= (+) i
let good1 = parseInt "1" >>= strAdd "2" >>= strAdd "3"
printf "%A" good1

// Outputs 6

let bad1 = parseInt "1" >>= strAdd "xyz" >>= strAdd "3"
printf "%A" bad1

// Outputs None

// Same as previous but using option wrapper
let maybeAdd (i:int option) v = 
    match v with
        | Some val1 -> Some(i.Value + val1)
        | None _ -> None

let maybeStrAdd str i = maybeParseInt str >>= maybeAdd i

let good2 = maybeParseInt "1" >>= maybeStrAdd "2" >>= maybeStrAdd "3"
printf "%A" good2

let bad2 = maybeParseInt "1" >>= maybeStrAdd "xyz" >>= maybeStrAdd "3"
printf "%A" bad2

