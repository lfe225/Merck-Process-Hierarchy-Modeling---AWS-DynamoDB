package main

import (
    "fmt"
    "os"
	"bufio" 

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
)

func GetOperationsByStage(tableName string) ([]Operation, error) {
    // Create a session with AWS
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("us-east-2"),
    })
    if err != nil {
        return nil, fmt.Errorf("failed to create session: %v", err)
    }

    // Create DynamoDB client
    svc := dynamodb.New(sess)

	//User prompt
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the Stage ID (e.g., 'Stage: Lehigh-1'): ")
	stageID, _ := reader.ReadString('\n')

	// Clean up the input
	stageID = stageID[:len(stageID)-1] 

    //Query Input
    input := &dynamodb.QueryInput{
        TableName:              aws.String(tableName),
        IndexName:              aws.String("GSI2"), 
        KeyConditionExpression:  aws.String("StageID = :stageID"),
        ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
            ":stageID": {
                S: aws.String(stageID),
            },
        },
    }

    //Query 
    result, err := svc.Query(input)
	//Error
    if err != nil {
        return nil, fmt.Errorf("failed to query operations: %v", err)
    }

    //Print out the operations 
    var operations []Operation
    for _, item := range result.Items {
        operation := Operation{
            Operation: *item["OperationID"].S,
        }
		fmt.Printf("Operation ID: %s\n", operation)
        operations = append(operations, operation)
    }

    return operations, nil
}

