 # 缓存


```python
# fn =warraper(fn)
# GetObject = cache(GetObject)
# GetObject = wrapper(Get(ctx,GetObject))
@cache(ttl=300)
def GetObject() {

}
```