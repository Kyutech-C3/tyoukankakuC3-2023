import { useState, ChangeEvent, useRef } from 'react'

interface ToDoItem {
  title: string
  detail: string
}

interface props {
  handleSubmit: any
}

const ToDoItem = (props: props) => {
  const [title, setTitle] = useState('')
  const [detail, setDetail] = useState('')

  const inputRef = useRef<HTMLInputElement>(null)

  const handleTitleChange = (event: ChangeEvent<HTMLInputElement>) => {
    setTitle(event.target.value)
  }

  const handleDetailChange = (event: ChangeEvent<HTMLInputElement>) => {
    setDetail(event.target.value)
  }

  const buttonClassHandler = () => {
    if (title !== '' && detail !== '') {
      return 'pointer-events-auto'
    } else {
      return 'pointer-events-none opacity-30'
    }
  }

  const keySubmitHandler = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.key === 'Enter' && (event.ctrlKey || event.metaKey)) {
      submit(title, detail)
    }
  }

  const submit = (t: string, d: string) => {
    if (t === '' || d === '') return
    props.handleSubmit(t, d)
    setTitle('')
    setDetail('')
    inputRef.current?.focus()
  }

  return (
    <div className="px-5 py-3 rounded-xl shadow-md shadow-gray-300 my-2 relative flex flex-col items-center">
      <label className="text-gray-600">
        タイトル :
        <input
          className="mx-3 px-2 py-0.5 bg-gray-100 rounded border-none focus:outline-none"
          type="text"
          name="title"
          placeholder="やること"
          value={title}
          ref={inputRef}
          onChange={handleTitleChange}
          onKeyDown={keySubmitHandler}
        />
      </label>
      <label className="text-gray-600">
        詳細 :
        <input
          className="mx-3 my-2 px-2 py-0.5 bg-gray-100 rounded border-none focus:outline-none"
          type="text"
          name="detail"
          placeholder="詳細"
          value={detail}
          onChange={handleDetailChange}
          onKeyDown={keySubmitHandler}
        />
      </label>
      <button
        className={`px-5 py-3 my-3 rounded-lg shadow-md w-fit select-none ${buttonClassHandler()}`}
        onClick={() => submit(title, detail)}
      >
        追加
      </button>
    </div>
  )
}

export default ToDoItem
