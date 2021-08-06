

#User: {
  FirstName: string
  LastName: string
  [=~"FirstName|LastName"]: =~"^[A-Z]{1}[a-zA-Z]{1,}"
}

{#User}