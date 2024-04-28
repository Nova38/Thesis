export default () => {
  const open = useState('SideBarOpen', () => true)

  const toggle = () => {
    open.value = !open.value
  }
}
