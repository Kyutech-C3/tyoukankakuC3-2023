interface BaseButtonProps {
  itemId: string
  text: string
  isActive: boolean
  handleDone: Function
}

const buttonActivityHandler = (isActive: boolean) => {
  if (isActive) {
    return 'pointer-events-auto bg-sky-500 hover:bg-sky-700 text-white'
  } else {
    return 'pointer-events-auto bg-gray-300 hover:bg-gray-400 text-black'
  }
}

const BaseButton = (props: BaseButtonProps) => {
  return (
    <div
      className={`px-3 py-0.5 text-xs w-fit rounded-md select-none cursor-pointer ${buttonActivityHandler(
        props.isActive
      )}`}
      onClick={() => props.handleDone(props.itemId)}
    >
      {props.text}
    </div>
  )
}

export default BaseButton
