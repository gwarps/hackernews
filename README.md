## Example queries

mutation {
  createLink(input:{title: "new link", address:"http://address.org"}) {
    title,
    address,
    id
  }
}


query{
  links{
    id
    title
    address
  }
}
