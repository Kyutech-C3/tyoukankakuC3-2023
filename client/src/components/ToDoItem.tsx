import { FaTrash } from 'react-icons/fa'
import dayjs from '../dayjs'

import BaseButton from './BaseButton'

interface ToDoItemProps {
  id: string
  title: string
  detail: string
  status: string
  created_at: string
  handleDelete: any
  handleDone: Function
}

const ToDoItem = (props: ToDoItemProps) => {
  return (
    <div className="px-5 pb-3 pt-10 rounded-xl shadow-md shadow-gray-300 my-2 relative">
      <div className="absolute top-3 right-5 flex items-center">
        <BaseButton
          itemId={props.id}
          text={props.status === 'done' ? 'ToDo' : 'Done'}
          isActive={props.status === 'todo'}
          handleDone={props.handleDone}
        />
        <div
          className="ml-3 text-gray-400 hover:text-gray-600"
          onClick={() => {
            props.handleDelete(props.id)
          }}
        >
          <FaTrash size="13px" />
        </div>
      </div>

      <div className="w-fit text-xl">{props.title}</div>
      <div className="w-fit text-gray-600 mx-2">{props.detail}</div>
      <div className="text-gray-500 text-right text-xs">
        {dayjs(props.created_at).tz('Asia/Tokyo').format('YYYY/MM/DD HH:mm:ss')}
      </div>
    </div>
  )
}

export default ToDoItem
