# Rule Engine UI

This project provides a simple AST rule engine. The frontend allows users to create rules, combine multiple rules, and evaluate rules against provided data. The backend serves the rule processing functionality in GoLang.

### Example of interaction flow:

1. Input your rule string in the "Create Rule" section and click `Create Rule`.
2. To combine multiple rules, enter one rule per line in the "Combine Rules" section and click `Combine Rules`.
3. In the "Evaluate Rule" section, input rule strings and details (in `key=value` format) to evaluate them. Click `Evaluate Rule` to get the result.

## Backend API Endpoints

The frontend interacts with the following backend API endpoints:

1. **Create Rule**
   - Endpoint: `POST /api/v1/create_rule`
   - Payload: `{ "rule_string": "your_rule_string" }`
   - Response: JSON object with the status of rule creation.

2. **Combine Rules**
   - Endpoint: `POST /api/v1/combine_rules`
   - Payload: `{ "rule_strings": ["rule1", "rule2"] }`
   - Response: JSON object with the combined rule result.

3. **Evaluate Rule**
   - Endpoint: `POST /api/v1/evaluate`
   - Payload: `{ "rules": ["rule1", "rule2"], "data": { "key": "value" } }`
   - Response: JSON object with the evaluation result.

## Dependencies
1. Docker
2. git

## Directions to run
1. Open directory in terminal
2. RUN commands
    - docker build -t my-go-app .
    - docker run -d -p 8000:8000 --name go-server my-go-app
3. Open browser visit http://localhost:8000/

## Sample Input
1. **Create Rule**
    age>21 OR (salary<=13000 AND salary>=10000)

2. **Combine Rules**
   age>21
   salary<=13000 AND salary>=10000
   department = Sales

4. **Evaluate Rule**
   Rule:
       age>21
       salary<=13000 AND salary>=10000
   Details:
       age = 22
       salary = 12000
