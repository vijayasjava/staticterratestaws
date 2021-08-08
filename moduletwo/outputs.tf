output "policydataoutput" {
  value = data.aws_iam_policy_document.ec2poweruser.json
}

output "userpolicydataoutputid" {
  value = aws_iam_user_policy.policy.id
}

output "userpolicydataoutputname" {
  value = aws_iam_user_policy.policy.name
}