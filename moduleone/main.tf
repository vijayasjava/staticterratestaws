resource "aws_iam_user" "user" {
  name = var.username
  path = "/system/"
}

resource "aws_iam_user_policy" "policy" {
  name = var.policyname
  user = aws_iam_user.user.name
  policy = data.aws_iam_policy_document.ec2poweruser.json
}