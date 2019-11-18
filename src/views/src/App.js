import React, { Component } from 'react';
import './components/employee';
import './App.css';
import Employee from './components/employee';

class App extends Component {
  
  constructor() {
    super();
    this.state = {
      employees: [],
      organisation: null,
    };
  }

  componentDidMount() {
    Promise.all([
      fetch('/api/v1/employees/'),
      fetch('/api/v1/employees/org')
    ])
    .then(([res1, res2]) => Promise.all([res1.json(), res2.json()]))
    .then(([data1, data2]) => {
        this.setState({ employees: data1, organisation: data2 })
    })
    .catch(console.log)
  }

  render() {
    return (
      <div className="App row">
        <div className="col-6">
          <h3>Employee Directory</h3>
          <table className="table table-striped table-responsive-md btn-table">
            <thead>
              <th>Employee Name</th>
              <th>Employee ID</th>
              <th>Manager ID</th>
            </thead>
            <tbody>
            {this.state.employees.map((e) => (
              <tr>
                <td>{e.name}</td>
                <td>{e.id}</td>
                <td>{e.managerId}</td>
              </tr>
            ))}
            </tbody>
          </table>
        </div> 
        <div className="col-6 list-container">
          <h3>Hierarchical View</h3>
          {this.state.organisation != null &&
            <div class="border">
            <h5 class="pt-3 pl-3">Managed Employees</h5>
            <hr/>
            { this.state.organisation['topmanagers'].map((e) => (
                 <Employee value={e}/> 
              ))
            }
            </div>
          }
          { this.state.organisation != null && this.state.organisation['freelancers'] != null &&
              <div class="border">
              <h5 class="pt-3 pl-3">Freelancers</h5>
              <hr/>
              { this.state.organisation['freelancers'].map((e) => (
                <Employee value={e}/> 
              ))}
              </div>
          }
           {this.state.organisation != null && this.state.organisation['managedByUnknown'] != null &&
              <div class="border">
              <h5 class="pt-3 pl-3">Managed By Unkown (Invalid Manager Id)</h5>
              <hr/>
              { this.state.organisation['managedByUnknown'].map((e) => (
                <Employee value={e}/> 
              ))}
              </div>
            }
            {this.state.organisation != null && this.state.organisation['circles'] != null &&
              <div class="border">
              <h5 class="pt-3 pl-3">Circles of Management (a ->b, b->c, c -> a)</h5>
              <hr/>
              { this.state.organisation['circles'].map((e) => (
                <div class="border">
                  {e.map((i) => (
                    <Employee value={i}/>
                  ))} 
                </div>
              ))}
              </div>
            }
        </div>
      </div>
    );
  }
}

export default App;
