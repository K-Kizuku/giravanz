# Redis用のサブネットグループ
resource "aws_elasticache_subnet_group" "redis_subnet_group" {
  name       = "redis-subnet-group"
  subnet_ids = [      aws_subnet.public1.id,
      aws_subnet.public2.id]  # 適切なサブネットIDに置き換え

  description = "Redis Subnet Group"
}

# Redis用のセキュリティグループ
resource "aws_security_group" "redis_sg" {
  name   = "redis-sg"
  vpc_id =  aws_vpc.main.id

  ingress {
    from_port   = 6379
    to_port     = 6379
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}


resource "aws_elasticache_replication_group" "example" {
  automatic_failover_enabled  = true
  preferred_cache_cluster_azs = ["ap-northeast-1a", "ap-northeast-1c"]
  replication_group_id        = "tf-rep-group-1"
  description                 = "example description"
  node_type                   = "cache.t2.micro"
  num_cache_clusters          = 2
  port                        = 6379
  subnet_group_name = aws_elasticache_subnet_group.redis_subnet_group.name
}