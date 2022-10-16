
type Address = Address of string
type Contact = {
    FirstName: string
    LastName: string
    Address: Address
}

let createContact first last =
    Contact { FirstName = first; LastName = last}

let updateAddress addr contact =
   { contact with Address = addr }

let createAddress addr: string =
    if addr.Length > 0 then
        Some (Address addr)
    else
        None

let c = createContact "Matheus" "Antunes" |> updateAddress (Address "Here")
printf "%a", c