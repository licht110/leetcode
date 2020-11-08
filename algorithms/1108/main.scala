object Solution {
    def defangIPaddr(address: String): String = {
        return address.replace(".", "[.]")
    }
}
