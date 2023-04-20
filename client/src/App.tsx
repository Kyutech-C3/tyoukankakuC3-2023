import { useState, useEffect } from 'react'
import ToDoItem from './components/ToDoItem'
import ToDoForm from './components/ToDoForm'
import axios from './axios'

interface ToDoItem {
  id: string
  title: string
  detail: string
  status: string
  created_at: string
  updated_at: string
}

function App() {
  const [toDoList, setToDoList] = useState<ToDoItem[]>([])
  const [doneList, setDoneList] = useState<ToDoItem[]>([])
  const [isSmallWindow, setIsSmallWindow] = useState<Boolean>()

  const fetchData = async () => {
    const response = await axios.get('/todo/')
    console.log(response.data)
    setDoneList([])
    setToDoList([])
    response.data.forEach((item: ToDoItem) => {
      if (item.status === 'done') {
        setDoneList((doneList) => [...doneList, item])
      } else if (item.status === 'todo') {
        setToDoList((toDoList) => [...toDoList, item])
      }
    })
    return response
  }

  const handleSubmit = async (title: string, detail: string) => {
    const response = await axios.post('/todo/', {
      title: title,
      detail: detail,
      status: 'todo',
    })
    console.log(response.data)
    fetchData()
  }

  const handleDelete = async (id: string) => {
    const response = await axios.delete(`/todo/${id}`)
    console.log(response.data)
    fetchData()
  }

  const handleDone = async (
    id: string,
    title: string,
    detail: string,
    status: string
  ) => {
    const response = await axios.patch(`/todo/${id}`, {
      title: title,
      detail: detail,
      status: status === 'todo' ? 'done' : 'todo',
    })
    console.log(response.data)
    fetchData()
  }

  const windowSizeHandler = () => {
    if (isSmallWindow) {
      return ''
    } else {
      return 'flex items-start justify-between'
    }
  }

  useEffect(() => {
    fetchData()
  }, [])

  window.addEventListener('resize', () => {
    if (window.innerWidth <= 570) {
      setIsSmallWindow(true)
    } else {
      setIsSmallWindow(false)
    }
  })

  return (
    <div className="todo-list max-w-[800px] min-w-[270px] w-[90%] mx-auto my-14">
      <div className="text-gray-600 text-4xl mb-10">ToDoアプリ</div>
      <ToDoForm
        handleSubmit={(title: string, detail: string) => {
          handleSubmit(title, detail)
        }}
      />

      <div className={`w-full mt-6 ${windowSizeHandler()}`}>
        <div className={isSmallWindow ? 'w-full mt-10' : 'w-[45%]'}>
          <div className="flex items-end justify-between mb-5">
            <p className="text-2xl text-gray-600">ToDo List</p>
            <p className="text-gray-400">{toDoList.length}件</p>
          </div>
          <div>
            {toDoList.map((item, index) => {
              return (
                <ToDoItem
                  key={index}
                  id={item.id}
                  title={item.title}
                  detail={item.detail}
                  status={item.status}
                  created_at={item.created_at}
                  handleDelete={(id: string) => {
                    handleDelete(id)
                  }}
                  handleDone={(id: string) => {
                    handleDone(id, item.title, item.detail, item.status)
                  }}
                />
              )
            })}
          </div>
        </div>

        <div className={isSmallWindow ? 'w-full mt-10' : 'w-[45%]'}>
          <div className="flex items-end justify-between mb-5">
            <p className="text-2xl text-gray-600">Done List</p>
            <p className="text-gray-400">{doneList.length}件</p>
          </div>
          <div className="opacity-50">
            {doneList.map((item, index) => {
              return (
                <ToDoItem
                  key={index}
                  id={item.id}
                  title={item.title}
                  detail={item.detail}
                  status={item.status}
                  created_at={item.created_at}
                  handleDelete={(id: string) => {
                    handleDelete(id)
                  }}
                  handleDone={(id: string) => {
                    handleDone(id, item.title, item.detail, item.status)
                  }}
                />
              )
            })}
          </div>
        </div>
      </div>
    </div>
  )
}

export default App
