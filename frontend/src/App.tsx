import {useState} from 'react';
//import {Greet} from "../wailsjs/go/main/App";
import {Button} from "./components/ui/button";
import {Input} from "./components/ui/input";
import {Plus, CheckCircle2, Circle, X} from "lucide-react";
import {cn} from "./lib/utils";


type Todo = {
    id: number
    text: string
    completed: boolean
}

function App() {
    const [todos, setTodos] = useState<Todo[]>([])
    const [newTodo, setNewTodo] = useState("")
    const [filter, setFilter] = useState<"all" | "active" | "completed">("all")
    

    const addTodo = (e: React.FormEvent) => {
        e.preventDefault()
        if (newTodo.trim() === "") return
    
        setTodos([
          ...todos,
          {
            id: Date.now(),
            text: newTodo.trim(),
            completed: false,
          },
        ])
        setNewTodo("")
      }
    
  const toggleTodo = (id: number) => {
    setTodos(todos.map((todo) => (todo.id === id ? { ...todo, completed: !todo.completed } : todo)))
  }

  const deleteTodo = (id: number) => {
    setTodos(todos.filter((todo) => todo.id !== id))
  }

  const completedCount = todos.filter((todo) => todo.completed).length
  const totalCount = todos.length

  const filteredTodos = todos.filter((todo) => {
    if (filter === "active") return !todo.completed
    if (filter === "completed") return todo.completed
    return true
  });

    return (
        <div className="max-w-md mx-auto p-4 space-y-4">
        <div className="space-y-4">
          <div>
            <h1 className="text-2xl font-bold">Todo List</h1>
          </div>
  
          <form onSubmit={addTodo} className="flex space-x-2">
            <Input
              type="text"
              placeholder="Add a new task..."
              value={newTodo}
              onChange={(e) => setNewTodo(e.target.value)}
              className="flex-1"
            />
            <Button type="submit">
              <Plus className="h-4 w-4 mr-2" />
              Add
            </Button>
          </form>
  
          <div className="flex items-center justify-between">
            <div className="text-sm text-muted-foreground">
              What to do ({completedCount}/{totalCount})
            </div>
            <div className="flex space-x-2">
              <Button variant={filter === "all" ? "default" : "outline"} size="sm" onClick={() => setFilter("all")}>
                All
              </Button>
              <Button variant={filter === "active" ? "default" : "outline"} size="sm" onClick={() => setFilter("active")}>
                Active
              </Button>
              <Button
                variant={filter === "completed" ? "default" : "outline"}
                size="sm"
                onClick={() => setFilter("completed")}
              >
                Completed
              </Button>
            </div>
          </div>
        </div>
  
        <ul className="space-y-2">
          {filteredTodos.map((todo) => (
            <li key={todo.id} className="flex items-center justify-between p-3 border rounded-md">
              <div className="flex items-center space-x-3">
                <Button variant="ghost" size="icon" onClick={() => toggleTodo(todo.id)} className="h-8 w-8">
                  {todo.completed ? <CheckCircle2 className="h-5 w-5 text-primary" /> : <Circle className="h-5 w-5" />}
                </Button>
                <span className={cn("text-sm", todo.completed && "line-through text-muted-foreground")}>{todo.text}</span>
              </div>
              <Button
                variant="ghost"
                size="icon"
                onClick={() => deleteTodo(todo.id)}
                className="h-8 w-8 text-destructive"
              >
                <X className="h-4 w-4" />
              </Button>
            </li>
          ))}
        </ul>
  
        {todos.length === 0 && <div className="text-center py-6 text-muted-foreground">No tasks yet. Add one above!</div>}
      </div>
    )
}

export default App
