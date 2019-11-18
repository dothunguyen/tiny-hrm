import React from 'react'

const Employee = ({value}) => {
  return (
    <ul  className="tree">
      <li>{value['name']}</li>
      {value['managerOf'] != null &&
        <li>
          <ul className="tree">
          {value['managerOf'].map((me) => (
              <li >
                <Employee value={me} />
              </li>
          ))}
          </ul>
        </li>
      }
    </ul>
  )
};

export default Employee