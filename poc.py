import requests as r
from datetime import datetime

url = "http://127.0.0.1:8000/"

def create():
    task_name = input("Input the task name... ")
    task_desc = input("Input task description... ")
    task_due_date = datetime.strptime(input("Input the due date (format: dd/mm/yy h:m)... "), '%d/%m/%y %H:%M').timestamp()
    response = r.put(url + "create", headers={"Content-Type": "application/json"}, json={
        "task_name": task_name,
        "task_desc": task_desc,
        "task_due": int(task_due_date),
        "task_mark": False
    })
    print(response.status_code, response.text)

def delete():
    ids = list(map(str, input("Input the ID of the records pending deletion... ").split()))
    response = r.delete(url + "delete_tasks", headers={"Content-Type": "application/json"}, json={
        "ids": ",".join(ids)
    })
    print(response.status_code, response.text)

def getAll():
    response = r.get(url, headers={"Content-Type": "application/json"})
    print(response.status_code, response.text)

def update():
    id = input("Input the ID of the record pending updater... ")
    response = r.post(url + "edit", headers={"Content-Type": "application/json"}, json={
        "task_id": id
    })
    print(response.status_code, response.text)

funcs = [getAll, create, delete, update, exit]
while True:
    func = funcs[int(input("Input the function number... ")) - 1]
    func()
    print("[getAll, create, delete, update, exit]")
