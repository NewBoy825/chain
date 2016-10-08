import React from 'react'

class ListItem extends React.Component {
  render() {
    const item = {...this.props.item}
    delete item.id

    return(
      <div className='panel panel-default'>
        <div className='panel-heading'>
          <strong>Key - {item.alias}</strong>
        </div>
        <div className='panel-body'>
          <pre>
            {JSON.stringify(item, null, '  ')}
          </pre>
        </div>
      </div>
    )
  }
}

export default ListItem