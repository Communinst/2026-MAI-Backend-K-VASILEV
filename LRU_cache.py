class LRUCache:
    def __init__(self, capacity: int=10):
        self.cap = capacity
        self.table = {}

    def get(self, key: str) -> str:
        if key not in self.table:
            return -1
        
        val = self.table.pop(key)
        self.table[key] = val
        return val

    def set(self, key: str, value: str) -> None:
        if key in self.table:
            del self.table[key]
            
        self.table[key] = value
        
        if len(self.table) > self.cap:
            oldest_key = next(iter(self.table))
            del self.table[oldest_key]

    def rem(self, key: str) -> None:
        self.table.pop(key, None)