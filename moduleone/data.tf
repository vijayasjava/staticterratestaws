data "aws_iam_policy_document" "ec2poweruser" {
 statement {
    actions = ["ec2:*"]
    effect = "Allow"
    resources = ["*"]
  }
}
