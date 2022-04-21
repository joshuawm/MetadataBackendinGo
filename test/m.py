import json

if __name__=="__main__":
    with open("test/data.json","r") as f:
        j = json.load(f)
        r=[]
        for x in j:
            i=x["Episode_ID"]
            r.append({
                "epmeta":{
                    "url":f"https://www.gevi.com/{i}",
                    "name":x["title"],
                    "desc":x["description"],
                    "series":"",
                    "runtime":0,
                    "tags":[],
                    
                },
                "name":x["studio"]
            })
        with open("result.json","w") as g :
            g.write(json.dumps(r))